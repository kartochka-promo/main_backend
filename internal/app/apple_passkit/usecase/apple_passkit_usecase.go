package usecase

import (
	"2020_1_drop_table/internal/app/apple_passkit"
	"2020_1_drop_table/internal/app/apple_passkit/models"
	"2020_1_drop_table/internal/app/cafe"
	cafeModels "2020_1_drop_table/internal/app/cafe/models"
	"2020_1_drop_table/internal/app/customer"
	customerModels "2020_1_drop_table/internal/app/customer/models"
	globalModels "2020_1_drop_table/internal/app/models"
	passesGenerator "2020_1_drop_table/internal/pkg/apple_pass_generator"
	"bytes"
	"context"
	"database/sql"
	"github.com/gorilla/sessions"
	"time"
)

type applePassKitUsecase struct {
	passKitRepo     apple_passkit.Repository
	cafeRepo        cafe.Repository
	customerRepo    customer.Repository
	passesGenerator passesGenerator.Generator
	contextTimeout  time.Duration
}

func NewApplePassKitUsecase(passKitRepo apple_passkit.Repository, cafeRepo cafe.Repository,
	customerRepo customer.Repository, passesGenerator passesGenerator.Generator,
	contextTimeout time.Duration) apple_passkit.Usecase {
	return &applePassKitUsecase{
		passKitRepo:     passKitRepo,
		cafeRepo:        cafeRepo,
		customerRepo:    customerRepo,
		passesGenerator: passesGenerator,
		contextTimeout:  contextTimeout,
	}
}

func passDBtoPassResource(db models.ApplePassDB) passesGenerator.ApplePass {
	files := map[string][]byte{
		"icon.png":    db.Icon,
		"icon@2x.png": db.Icon2x,
		"logo.png":    db.Logo,
		"logo@2x.png": db.Logo2x,
	}

	if len(db.Strip) != 0 && len(db.Strip2x) != 0 {
		files["strip.png"] = db.Strip
		files["strip@2x.png"] = db.Strip2x
	}

	return passesGenerator.NewApplePass(db.Design, files)
}

func (ap *applePassKitUsecase) addNewSavedPassToCafe(ctx context.Context, pass models.ApplePassDB,
	cafeObj cafeModels.Cafe) error {

	newPass, err := ap.passKitRepo.Add(ctx, pass)
	if err != nil {
		return err
	}

	newPassId := sql.NullInt64{
		Int64: int64(newPass.ApplePassID),
		Valid: true,
	}
	cafeObj.SavedApplePassID = newPassId
	_ = ap.cafeRepo.UpdateSavedPass(ctx, cafeObj)

	return nil
}

func (ap *applePassKitUsecase) addNewPublishedPassToCafe(ctx context.Context, pass models.ApplePassDB,
	cafeObj cafeModels.Cafe) error {

	newPass, err := ap.passKitRepo.Add(ctx, pass)
	if err != nil {
		return err
	}

	newPassId := sql.NullInt64{
		Int64: int64(newPass.ApplePassID),
		Valid: true,
	}
	cafeObj.PublishedApplePassID = newPassId
	_ = ap.cafeRepo.UpdatePublishedPass(ctx, cafeObj)

	return nil
}

func (ap *applePassKitUsecase) UpdatePass(c context.Context, pass models.ApplePassDB, cafeID int, publish bool) error {
	ctx, cancel := context.WithTimeout(c, ap.contextTimeout)
	defer cancel()

	session := ctx.Value("session").(*sessions.Session)

	staffInterface, found := session.Values["userID"]
	staffID, ok := staffInterface.(int)

	if !found || !ok || staffID <= 0 {
		return globalModels.ErrForbidden
	}

	cafeObj, err := ap.cafeRepo.GetByID(ctx, cafeID)
	if err != nil {
		return err
	}

	if cafeObj.StaffID != staffID {
		return globalModels.ErrForbidden
	}

	if cafeObj.SavedApplePassID.Valid {
		pass.ApplePassID = int(cafeObj.SavedApplePassID.Int64)
		err := ap.passKitRepo.Update(ctx, pass)
		if err != nil {
			return err
		}
	} else {
		err = ap.addNewSavedPassToCafe(ctx, pass, cafeObj)
		if err != nil {
			return err
		}
	}

	if !publish {
		return nil
	}

	if cafeObj.PublishedApplePassID.Valid {
		pass.ApplePassID = int(cafeObj.PublishedApplePassID.Int64)
		err := ap.passKitRepo.Update(ctx, pass)
		if err != nil {
			return err
		}
	} else {
		err = ap.addNewPublishedPassToCafe(ctx, pass, cafeObj)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ap *applePassKitUsecase) GeneratePassObject(c context.Context, cafeID int) (*bytes.Buffer, error) {
	//ToDo make not only for published cards
	ctx, cancel := context.WithTimeout(c, ap.contextTimeout)
	defer cancel()

	newCustomer := customerModels.Customer{CafeID: cafeID}

	newCustomer, err := ap.customerRepo.Add(ctx, newCustomer)
	if err != nil {
		return nil, err
	}

	//ToDo make valid barcode generation

	cafeObj, err := ap.cafeRepo.GetByID(ctx, cafeID)
	if err != nil {
		return nil, err
	}
	if !cafeObj.PublishedApplePassID.Valid {
		return nil, globalModels.ErrNoPublishedCard
	}

	publishedCardID := int(cafeObj.PublishedApplePassID.Int64)
	publishedCardDB, err := ap.passKitRepo.GetByID(ctx, publishedCardID)
	if err != nil {
		return nil, err
	}
	passBuffer, err := ap.passesGenerator.CreateNewPass(passDBtoPassResource(publishedCardDB))

	return passBuffer, err
}