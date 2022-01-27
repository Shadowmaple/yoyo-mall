package service

import "yoyo-mall/service/user"

type Service struct {
	User *user.Service
}

func (s *Service) Init() {

}
