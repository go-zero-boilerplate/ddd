package default_encryption_svc

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Authentication Testing", t, func() {

		Convey("generateSalt()", func() {
			salt, err := generateSalt(cDEFAULT_SALT_LENGTH)
			So(err, ShouldBeNil)
			So(len(salt), ShouldBeGreaterThan, 0)
			So(len(salt), ShouldEqual, cDEFAULT_SALT_LENGTH)
		})

		Convey("combine()", func() {
			salt, err := generateSalt(cDEFAULT_SALT_LENGTH)
			So(err, ShouldBeNil)
			password := "boomchuckalucka"
			expectedLength := len(salt) + len(password)
			combo := combine(salt, []byte(password))

			So(len(combo), ShouldBeGreaterThan, 0)
			So(len(combo), ShouldEqual, expectedLength)
			So(bytes.Equal(combo[len(combo)-len(salt):], salt), ShouldBeTrue)
		})

		Convey("hashPassword()", func() {
			salt, err := generateSalt(cDEFAULT_SALT_LENGTH)
			So(err, ShouldBeNil)
			combo := combine(salt, []byte("hershmahgersh"))
			hash, err := hashPassword(combo, CDEFAULT_ENCRYPT_COST)
			So(err, ShouldBeNil)
			So(len(hash), ShouldBeGreaterThan, 0)

			cost, err := bcrypt.Cost([]byte(hash))
			if err != nil {
				log.Print(err)
			}
			So(cost, ShouldEqual, CDEFAULT_ENCRYPT_COST)
		})

		Convey("CreatePassword()", func() {
			passString := "mmmPassword1"
			pwd, err := CreatePassword(passString, cDEFAULT_SALT_LENGTH, CDEFAULT_ENCRYPT_COST)
			So(err, ShouldBeNil)
			pass_struct := new(password)

			So(pwd, ShouldHaveSameTypeAs, pass_struct)
			So(len(pwd.hash), ShouldBeGreaterThan, 0)
			So(len(pwd.salt), ShouldBeGreaterThan, 0)
			So(len(pwd.salt), ShouldEqual, cDEFAULT_SALT_LENGTH)
		})

		Convey("comparePassword", func() {
			password := "megaman49"
			passwordMeta, err := CreatePassword(password, cDEFAULT_SALT_LENGTH, CDEFAULT_ENCRYPT_COST)
			So(err, ShouldBeNil)

			So(passwordMatch(password, passwordMeta.salt, passwordMeta.hash), ShouldBeTrue)
			So(passwordMatch("lolfail", passwordMeta.salt, passwordMeta.hash), ShouldBeFalse)
			So(passwordMatch("Megaman49", passwordMeta.salt, passwordMeta.hash), ShouldBeFalse)
		})
	})
}
