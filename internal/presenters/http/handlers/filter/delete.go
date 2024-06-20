package filter

import "github.com/labstack/echo/v4"

// DeleteFilter godoc
// @Summary Delete Filter
// @Description Delete Filter for ProjectRPCPayload by projectId. Roles allowed: business, staff
// @Tags Filter
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query int true "filter id"
// @Success 204 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/filters/{filterId} [delete]
func (fh Handler) DeleteFilter(c echo.Context) error {
	return nil
	//accessTokenClaims, _ := c.Get("accessTokenClaims").(*OAuthService.AccessTokenClaims)
	//filterId, err := strconv.Atoi(c.Param("filterId"))
	//if err != nil {
	//	return h.ErrorResponse(c, http.StatusBadRequest, "filter id does not valid", err)
	//}
	//
	//filter, err := h.filterRepo.FindById(filterId)
	//if err != nil {
	//	return h.ErrorResponse(c, http.StatusInternalServerError, "could not find filter", err)
	//}
	//if filter == nil {
	//	return h.ErrorResponse(c, http.StatusNotFound, "filter not found", nil)
	//}
	//
	//switch accessTokenClaims.UserRole {
	//case "business":
	//	result, err := RpcClient.GetProject(filter.ProjectToken)
	//	if err != nil {
	//		return h.ErrorResponse(c, http.StatusInternalServerError, "could not delete filter", err)
	//	}
	//	if result.Token != filter.ProjectToken {
	//		return h.ErrorResponse(c, http.StatusForbidden, "access denied", err)
	//	}
	//case "staff":
	//	break
	//default:
	//	return h.ErrorResponse(c, http.StatusForbidden, "access denied", nil)
	//}
	//if err := h.filterRepo.Delete(filter); err != nil {
	//	return h.ErrorResponse(c, http.StatusInternalServerError, "could not delete filter", err)
	//}
	//return h.SuccessResponse(c, http.StatusNoContent, "filter successfully deleted", nil)
}
