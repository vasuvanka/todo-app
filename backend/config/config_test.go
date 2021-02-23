package config_test

import (
	"testing"

	"github.com/vasuvanka/todo-app/backend/config"
)

func TestNew(t *testing.T)  {
	conf := config.New()
	if conf.DatabaseURI != "" {
		t.Errorf("got %s, want empty string", conf.DatabaseURI)
	}
	if conf.Port != "" {
		t.Errorf("got %s, want empty string", conf.Port)
	}
	if conf.Env != "" {
		t.Errorf("got %s, want empty string", conf.Env)
	}
	if conf.FEPath != "" {
		t.Errorf("got %s, want empty string", conf.FEPath)
	}
}

func TestInit(t *testing.T){
	conf := config.New();
	conf.Init()
	if conf.DatabaseURI != "mongodb://localhost:27017/todo" {
		t.Errorf("got %s, want mongodb://localhost:27017/todo", conf.DatabaseURI)
	}
	if conf.Port != "8080" {
		t.Errorf("got %s, want 8080", conf.Port)
	}
	if conf.Env != "dev" {
		t.Errorf("got %s, want dev", conf.Env)
	}
	if conf.FEPath != "./dist/todo" {
		t.Errorf("got %s, want ./dist/todo", conf.FEPath)
	}
}