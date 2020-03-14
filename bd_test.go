package main

import (
	"2020_1_drop_table/owners"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAppend(t *testing.T) {
	Storage, _ := owners.NewOwnerStorage("postgres", "", "5431")
	Storage.Clear()
	own := owners.Owner{
		Name:     "asd",
		Email:    "asd",
		Password: "asd",
		EditedAt: time.Now(),
		Photo:    "asd",
	}
	_, err := Storage.Append(own)
	fmt.Println(err)
	assert.Nil(t, err, "No errors")
	own2 := owners.Owner{
		Name:     "asd",
		Email:    "assd",
		Password: "asd",
		EditedAt: time.Now(),
		Photo:    "asd",
	}
	_, err2 := Storage.Append(own2)

	assert.Nil(t, err2, "No erros")
}

func TestGet(t *testing.T) {
	Storage, _ := owners.NewOwnerStorage("postgres", "", "5431")
	Storage.Clear()
	own := owners.Owner{
		OwnerID:  1,
		Name:     "asd",
		Email:    "asd",
		Password: "asd",
		EditedAt: time.Now().UTC(),
		Photo:    "asd",
	}
	_, _ = Storage.Append(own)
	dbOwner, err := Storage.Get(1)
	assert.Nil(t, err)
	own.Password = owners.GetMD5Hash(own.Password)
	assert.Equal(t, own, dbOwner)
}

func TestOwnerStorage_GetByEmailAndPassword(t *testing.T) {
	Storage, _ := owners.NewOwnerStorage("postgres", "", "5431")
	Storage.Clear()
	own := owners.Owner{
		OwnerID:  1,
		Name:     "asd",
		Email:    "email",
		Password: "password",
		EditedAt: time.Now().UTC(),
		Photo:    "asd",
	}
	_, _ = Storage.Append(own)
	own.Password = owners.GetMD5Hash(own.Password)
	dbOwner, err := Storage.GetByEmailAndPassword("email", "password")
	assert.Nil(t, err)
	assert.Equal(t, own, dbOwner)
}

func TestOwnerStorage_Set(t *testing.T) {
	Storage, _ := owners.NewOwnerStorage("postgres", "", "5431")
	Storage.Clear()
	own := owners.Owner{
		OwnerID:  1,
		Name:     "asd",
		Email:    "email",
		Password: "password",
		EditedAt: time.Now().UTC(),
		Photo:    "asd",
	}

	newOwn := owners.Owner{
		OwnerID:  1,
		Name:     "newasd",
		Email:    "newemail",
		Password: "password",
		EditedAt: time.Now().UTC(),
		Photo:    "asd",
	}
	_, _ = Storage.Append(own)
	_, _ = Storage.Set(1, newOwn)
	dBOwner, err := Storage.Get(1)
	assert.Nil(t, err)
	newOwn.Password = owners.GetMD5Hash(newOwn.Password)
	assert.Equal(t, newOwn, dBOwner)
}

func TestOwnerStorage_Count(t *testing.T) {
	Storage, _ := owners.NewOwnerStorage("postgres", "", "5431")
	Storage.Clear()
	count, err := Storage.Count()
	assert.Nil(t, err)
	assert.Equal(t, 0, count)

	own := owners.Owner{
		OwnerID:  229,
		Name:     "asd",
		Email:    "email",
		Password: "password",
		EditedAt: time.Now().UTC(),
		Photo:    "asd",
	}
	_, _ = Storage.Append(own)

	count, err = Storage.Count()
	assert.Nil(t, err)
	assert.Equal(t, 1, count)

}

func TestOwnerStorage_Existed(t *testing.T) {
	Storage, _ := owners.NewOwnerStorage("postgres", "", "5431")
	Storage.Clear()
	isExist, _, err := Storage.Existed("email", "password")
	assert.Nil(t, err)
	assert.Equal(t, false, isExist)

	own := owners.Owner{
		OwnerID:  229,
		Name:     "asd",
		Email:    "email",
		Password: "password",
		EditedAt: time.Now().UTC(),
		Photo:    "asd",
	}
	_, _ = Storage.Append(own)

	isExist, _, err = Storage.Existed("email", "password")
	assert.Nil(t, err)
	assert.Equal(t, true, isExist)
}
