package mgm_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wshops/mgm"
	"github.com/wshops/mgm/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

const mongo_uri = ""

// Note: to run Transaction tests, the MongoDB daemon must run as replica set, not as a standalone daemon.
// To convert it [see this](https://docs.mongodb.com/manual/tutorial/convert-standalone-to-replica-set/)
func TestTransactionCommit(t *testing.T) {
	util.PanicErr(
		mgm.SetDefaultConfig(nil, "models", options.Client().ApplyURI(mongo_uri)),
	)
	resetCollection()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.Transaction(func(session mongo.Session, sc mongo.SessionContext) error {

		err := mgm.Coll(d).CreateWithCtx(sc, d)

		if err != nil {
			return err
		}

		return session.CommitTransaction(sc)
	})

	util.AssertErrIsNil(t, err)
	count, err := mgm.Coll(d).CountDocuments(mgm.Ctx(), bson.M{})

	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(1), count)
}

func TestTransactionAbort(t *testing.T) {
	util.PanicErr(
		mgm.SetDefaultConfig(nil, "models", options.Client().ApplyURI(mongo_uri)),
	)
	resetCollection()
	//seed()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.Transaction(func(session mongo.Session, sc mongo.SessionContext) error {

		err := mgm.Coll(d).CreateWithCtx(sc, d)

		if err != nil {
			return err
		}

		return session.AbortTransaction(sc)
	})

	util.AssertErrIsNil(t, err)
	count, err := mgm.Coll(d).CountDocuments(mgm.Ctx(), bson.M{})

	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(0), count)
}

func TestTransactionWithCtx(t *testing.T) {
	util.PanicErr(
		mgm.SetDefaultConfig(nil, "models", options.Client().ApplyURI(mongo_uri)),
	)
	resetCollection()
	//seed()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.TransactionWithCtx(mgm.Ctx(), func(session mongo.Session, sc mongo.SessionContext) error {

		err := mgm.Coll(d).CreateWithCtx(sc, d)

		if err != nil {
			return err
		}

		return session.AbortTransaction(sc)
	})

	util.AssertErrIsNil(t, err)
	count, err := mgm.Coll(d).CountDocuments(mgm.Ctx(), bson.M{})

	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(0), count)
}
