package admin

import (
	"context"
	"fmt"
	"strconv"

	"github.com/goli-nababa/golibaba-backend/common"
	user "github.com/goli-nababa/golibaba-backend/modules/user_service_client"
)

type AdminService struct {
	userClient user.UserServiceClient
	//companyClient      company.CompanyServiceClient
	//hotelClient        hotel.HotelServiceClient
	//travelAgencyClient travelAgency.TravelAgencyServiceClient
}

func NewAdminService(userClient user.UserServiceClient,

// companyClient company.CompanyServiceClient,
// hotelClient hotel.HotelServiceClient,
// travelAgencyClient travelAgency.TravelAgencyServiceClient
) *AdminService {
	return &AdminService{
		userClient: userClient,
		//companyClient:      companyClient,
		//hotelClient:        hotelClient,
		//travelAgencyClient: travelAgencyClient,
	}
}

func (s *AdminService) BlockEntity(ctx context.Context, entityID, entityType string) error {
	switch entityType {
	case "user":
		userID, err := strconv.ParseUint(entityID, 10, 64) // Convert string to uint64
		if err != nil {
			return fmt.Errorf("invalid user ID: %w", err)
		}

		err = s.userClient.BlockUser(common.UserID(userID))
		if err != nil {
			return fmt.Errorf("failed to block user: %w", err)
		}

		err = s.BlockUserProperties(ctx, userID)
		if err != nil {
			return fmt.Errorf("failed to block properties for user %d: %w", userID, err)
		}
	/*case "company":
		_, err := s.companyClient.BlockCompany(ctx, entityID) //implement me
		return err
	case "hotel":
		_, err := s.hotelClient.BlockHotel(ctx, entityID) //implement me
		return err
	case "travel_agency":
		_, err := s.travelAgencyClient.BlockTravelAgency(ctx, entityID) //implement me
		return err*/
	default:
		return fmt.Errorf("unsupported entity type: %s", entityType)
	}

	return nil
}

func (s *AdminService) UnblockEntity(ctx context.Context, entityID, entityType string) error {
	switch entityType {
	case "user":
		userID, err := strconv.ParseUint(entityID, 10, 64) // Convert string to uint64
		if err != nil {
			return fmt.Errorf("invalid user ID: %w", err)
		}

		err = s.userClient.UnblockUser(common.UserID(userID))
		if err != nil {
			return fmt.Errorf("failed to unblock user: %w", err)
		}

		err = s.UnblockUserProperties(ctx, userID)
		if err != nil {
			return fmt.Errorf("failed to block properties for user %d: %w", userID, err)
		}
	/*case "company":
		_, err := s.companyClient.UnblockCompany(ctx, entityID)
		return err
	case "hotel":
		_, err := s.hotelClient.UnblockHotel(ctx, entityID)
		return err
	case "travel_agency":
		_, err := s.travelAgencyClient.UnblockTravelAgency(ctx, entityID)
		return err*/
	default:
		return fmt.Errorf("unsupported entity type: %s", entityType)
	}

	return nil
}

func (s *AdminService) AssignRole(ctx context.Context, userID uint, role string) error {
	return s.userClient.AssignRole(common.UserID(userID), role)
}

func (s *AdminService) CancelRole(ctx context.Context, userID uint, role string) error {
	return s.userClient.CancelRole(common.UserID(userID), role)
}

func (s *AdminService) AssignPermissionToRole(ctx context.Context, userID uint, role string, permissions []string) error {
	return s.userClient.AssignPermissionToRole(common.UserID(userID), role, permissions)
}
func (s *AdminService) RevokePermissionFromRole(ctx context.Context, userID uint, role string, permissions []string) error {
	return s.userClient.RevokePermissionFromRole(common.UserID(userID), role, permissions)
}
func (s *AdminService) PublishStatement(ctx context.Context, userIDs []common.UserID, action string, permissions []string) error {
	switch action {
	case "allow":
		return s.userClient.PublishStatement(userIDs, common.StatementActionAllow, permissions)
	case "deny":
		return s.userClient.PublishStatement(userIDs, common.StatementActionDeny, permissions)
	default:
		return fmt.Errorf("invalid action: %s", action)
	}
}
func (s *AdminService) CancelStatement(ctx context.Context, userID uint, statementID uint) error {
	return s.userClient.CancelStatement(common.UserID(userID), common.StatementID(statementID))
}

func (s *AdminService) BlockUserProperties(ctx context.Context, userID uint64) error {
	/*companies, err := s.companyClient.GetUserCompanies(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get companies for user %s: %w", userID, err)
	}

	for _, company := range companies {
		_, err := s.companyClient.BlockCompany(ctx, company.ID)
		if err != nil {
			return fmt.Errorf("failed to block company %s: %w", company.ID, err)
		}
	}

	hotels, err := s.hotelClient.GetUserHotels(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get hotels for user %s: %w", userID, err)
	}

	for _, hotel := range hotels {
		_, err := s.hotelClient.BlockHotel(ctx, hotel.ID)
		if err != nil {
			return fmt.Errorf("failed to block hotel %s: %w", hotel.ID, err)
		}
	}

	travelAgencies, err := s.travelAgencyClient.GetUserTravelAgencies(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get travel agencies for user %s: %w", userID, err)
	}

	for _, agency := range travelAgencies {
		_, err := s.travelAgencyClient.BlockTravelAgency(ctx, agency.ID)
		if err != nil {
			return fmt.Errorf("failed to block travel agency %s: %w", agency.ID, err)
		}
	}*/

	return nil
}

func (s *AdminService) UnblockUserProperties(ctx context.Context, userID uint64) error {
	/*companies, err := s.companyClient.GetUserCompanies(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get companies for user %s: %w", userID, err)
	}

	for _, company := range companies {
		_, err := s.companyClient.UnblockCompany(ctx, company.ID)
		if err != nil {
			return fmt.Errorf("failed to unblock company %s: %w", company.ID, err)
		}
	}

	hotels, err := s.hotelClient.GetUserHotels(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get hotels for user %s: %w", userID, err)
	}

	for _, hotel := range hotels {
		_, err := s.hotelClient.UnblockHotel(ctx, hotel.ID)
		if err != nil {
			return fmt.Errorf("failed to unblock hotel %s: %w", hotel.ID, err)
		}
	}

	travelAgencies, err := s.travelAgencyClient.GetUserTravelAgencies(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get travel agencies for user %s: %w", userID, err)
	}

	for _, agency := range travelAgencies {
		_, err := s.travelAgencyClient.UnblockTravelAgency(ctx, agency.ID)
		if err != nil {
			return fmt.Errorf("failed to unblock travel agency %s: %w", agency.ID, err)
		}
	}*/

	return nil
}
