package http

import (
	"context"
	"knowledgemap_backend/agent/knowledgemap/server/http/comm"
	"knowledgemap_backend/microservices/common/middlewares"
	capi "knowledgemap_backend/microservices/knowledgemap/class/api"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func classCreate(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(capi.CreateClassReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	/*
		没传教师信息，说明是教师自己创建
		否则，是教秘创建
	*/
	if req.Teachername == "" {
		req.Teacherid = c.Request().Header.Get("auth-uid")
		req.Teachername = c.Get("userName").(string)
	}
	if res, err := classSrv.CreateClass(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func joinClass(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(capi.JoinClassReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	req.Userid = c.Request().Header.Get("auth-uid")
	req.Username = c.Get("userName").(string)
	if res, err := classSrv.JoinClass(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryMyClasses(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(uapi.UserReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	req.Userid = c.Request().Header.Get("auth-uid")
	if res, err := classSrv.UserClassInfo(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryAllUserInClass(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(capi.ClassReq)
	req.Classid = c.Param("classid")
	if res, err := classSrv.QueryClassUserInfo(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}

func queryClassInfo(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(capi.ClassReq)
	req.Classid = c.Param("classid")
	if res, err := classSrv.ClassInfo(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}

func searchClasses(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(capi.SearchClassesInfoReq)
	req.Course = c.Param("course")
	req.College = c.Param("college")
	req.Subject = c.Param("subject")
	req.Page, _ = strconv.ParseInt(c.Param("page"), 10, 64)

	if res, err := classSrv.SearchClassesInfo(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	return c.JSON(http.StatusOK, comm.Data(nil))
}

func createInvitation(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(capi.InvitationReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := classSrv.CreateInvitaion(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func dropInvitation(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(capi.InvitationReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	if res, err := classSrv.StopInvitaion(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryInvitation(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(capi.InvitationReq)
	req.Invitaioncode = c.Param("invitationcode")
	if res, err := classSrv.InvitaionInfo(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func queryFormList(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(uapi.Empty)
	if res, err := classSrv.QueryFormList(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}
