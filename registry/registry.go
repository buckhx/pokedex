package registry

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sync"
	"time"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/registry/mint"
	"github.com/buckhx/safari-zone/util"
	"github.com/buckhx/safari-zone/util/kvc"
)

var validator = regexp.MustCompile(`^[a-zA-Z0-9]{3,32}$`)

const (
	Issuer   = "buckhx.safari.registry"
	TokenDur = 24 * time.Hour
)

type Registry struct {
	sync.Mutex
	db   kvc.KVC
	mint *mint.Mint
}

// pem is the path to .pem private key used to sign tokens
func New(pemfile string) (r *Registry, err error) {
	k, err := ioutil.ReadFile(pemfile)
	if err != nil {
		return
	}
	m, err := mint.NewEC256(Issuer, k)
	if err != nil {
		return
	}
	r = &Registry{
		db:   kvc.NewMem(),
		mint: m,
	}
	return
}

func (r *Registry) Add(req *pbf.Trainer) (err error) {
	switch {
	case !validator.MatchString(req.Name):
		err = fmt.Errorf("User name must match /%s/", validator)
	case !validator.MatchString(req.Password):
		err = fmt.Errorf("Password must match /%s/", validator)
	case req.Age < 10:
		err = fmt.Errorf("Trainer is too young!")
	case req.Age > 99:
		err = fmt.Errorf("Trainer is too old!")
	}
	if err != nil {
		return
	}
	req.Password = util.Hash(req.Password)
	req.Start = &pbf.Timestamp{Unix: time.Now().Unix()}
	req.Pc = &pbf.Pokemon_Collection{}
	r.Lock() // for race w/ GenUID
	defer r.Unlock()
	uid := util.GenUID()
	for r.db.Has(uid) {
		uid = util.GenUID()
	}
	req.Uid = uid
	r.db.Set(uid, req)
	return
}

func (r *Registry) Get(uid string) (*pbf.Trainer, error) {
	v := r.db.Get(uid)
	if v == nil {
		return nil, fmt.Errorf("Invalid trainer: Not registered")
	}
	if u, ok := v.(*pbf.Trainer); !ok {
		return nil, fmt.Errorf("Internal error: DB Assertion")
	} else {
		return u, nil
	}
}

func (r *Registry) Authenticate(req *pbf.Trainer) (tok *pbf.Token, err error) {
	v, err := r.Get(req.Uid)
	switch {
	case err != nil:
		break
	case v.Password == util.Hash(req.Password):
		err = fmt.Errorf("Invalid trainer: Password")
	case user{v}.hasScope(req.Scope...):
		err = fmt.Errorf("Invalid trainer: Scope")
	}
	if err != nil {
		return
	}
	if sig, err := r.mint.IssueToken(req.Uid, TokenDur, req.Scope...); err == nil {
		tok = &pbf.Token{Access: sig, Scope: req.Scope}
	}
	return
}