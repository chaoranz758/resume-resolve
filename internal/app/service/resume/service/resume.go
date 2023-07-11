package service

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"resume-resolving/api/idl/service/post/kitex_gen/post"
	"resume-resolving/api/idl/service/resume/kitex_gen/base"
	"resume-resolving/api/idl/service/resume/kitex_gen/resume"
	resume1 "resume-resolving/internal/app/service/resume"
	"resume-resolving/internal/app/service/resume/model"
	"resume-resolving/internal/app/service/resume/pkg/code"
	"resume-resolving/internal/pkg"
	"time"
)

const (
	tableLength = 7
	maxLength   = 10
	breakLength = 9
)

func UploadStructResume(request *resume.UploadStructResumeRPCRequest) (int32, string, error) {
	resumeId := resume1.GlobalEngine.Options.Id.GenId()

	value := make([][]interface{}, 0, tableLength)
	length := 1

	var basicInfo = model.BasicInfo{
		ResumeId:       resumeId,
		UserId:         request.BasicInformation.UserId,
		Name:           request.BasicInformation.Name,
		Phone:          request.BasicInformation.Phone,
		ResumeUrl:      request.BasicInformation.ResumeUrl,
		Email:          request.BasicInformation.Email,
		SelfEvaluation: request.BasicInformation.SelfEvaluaton,
		Birthday:       time.UnixMilli(request.BasicInformation.Birthday),
	}

	var basicInfoValue = []interface{}{&basicInfo}
	value = append(value, basicInfoValue)

	educationalExperienceList := make([]*model.EducationalExperience, 0, len(request.EducationalExperienceList))
	if len(request.EducationalExperienceList) != 0 {
		for i := 0; i < len(request.EducationalExperienceList); i++ {
			var educationalExperience = model.EducationalExperience{
				Id:         resume1.GlobalEngine.Options.Id.GenId(),
				ResumeId:   resumeId,
				School:     request.EducationalExperienceList[i].School,
				Education:  request.EducationalExperienceList[i].Education,
				Speciality: request.EducationalExperienceList[i].Speciality,
				Ranking:    request.EducationalExperienceList[i].Rank,
				StartTime:  time.UnixMilli(request.EducationalExperienceList[i].StartTime),
				EndTime:    time.UnixMilli(request.EducationalExperienceList[i].EndTime),
			}
			educationalExperienceList = append(educationalExperienceList, &educationalExperience)
		}

		var educationalExperienceListValue = []interface{}{educationalExperienceList}
		value = append(value, educationalExperienceListValue)
		length++
	}

	internshipExperienceList := make([]*model.InternshipWorkExperience, 0, len(request.InternshipExperienceList)+len(request.WorkExperienceList))
	if len(request.InternshipExperienceList) != 0 || len(request.WorkExperienceList) != 0 {
		for i := 0; i < len(request.InternshipExperienceList); i++ {
			var internshipExperience = model.InternshipWorkExperience{
				Id:          resume1.GlobalEngine.Options.Id.GenId(),
				ResumeId:    resumeId,
				Company:     request.InternshipExperienceList[i].Company,
				Position:    request.InternshipExperienceList[i].Position,
				Description: request.InternshipExperienceList[i].Description,
				StartTime:   time.UnixMilli(request.InternshipExperienceList[i].StartTime),
				EndTime:     time.UnixMilli(request.InternshipExperienceList[i].EndTime),
			}
			internshipExperienceList = append(internshipExperienceList, &internshipExperience)
		}
		for i := 0; i < len(request.WorkExperienceList); i++ {
			var workExperience = model.InternshipWorkExperience{
				Id:           resume1.GlobalEngine.Options.Id.GenId(),
				ResumeId:     resumeId,
				Company:      request.WorkExperienceList[i].Company,
				Position:     request.WorkExperienceList[i].Position,
				Description:  request.WorkExperienceList[i].Description,
				IsInternship: 1,
				StartTime:    time.UnixMilli(request.WorkExperienceList[i].StartTime),
				EndTime:      time.UnixMilli(request.WorkExperienceList[i].EndTime),
			}
			internshipExperienceList = append(internshipExperienceList, &workExperience)
		}
		var internshipExperienceListValue = []interface{}{internshipExperienceList}
		value = append(value, internshipExperienceListValue)
		length++
	}

	projectExperienceList := make([]*model.ProjectExperience, 0, len(request.ProjectExperienceList))
	if len(request.ProjectExperienceList) != 0 {
		for i := 0; i < len(request.ProjectExperienceList); i++ {
			var projectExperience = model.ProjectExperience{
				Id:                 resume1.GlobalEngine.Options.Id.GenId(),
				ResumeId:           resumeId,
				ProjectName:        request.ProjectExperienceList[i].ProjectName,
				ProjectRole:        request.ProjectExperienceList[i].ProjectRole,
				ProjectDescription: request.ProjectExperienceList[i].ProjectDescription,
				ProjectUrl:         request.ProjectExperienceList[i].ProjectUrl,
				StartTime:          time.UnixMilli(request.ProjectExperienceList[i].StartTime),
				EndTime:            time.UnixMilli(request.ProjectExperienceList[i].EndTime),
			}
			projectExperienceList = append(projectExperienceList, &projectExperience)
		}
		var projectExperienceListValue = []interface{}{projectExperienceList}
		value = append(value, projectExperienceListValue)
		length++
	}

	contestList := make([]*model.ContestCertificate, 0, len(request.ContestList)+len(request.CertificateList))
	if len(request.ContestList) != 0 || len(request.CertificateList) != 0 {
		for i := 0; i < len(request.ContestList); i++ {
			var contest = model.ContestCertificate{
				Id:          resume1.GlobalEngine.Options.Id.GenId(),
				ResumeId:    resumeId,
				Name:        request.ContestList[i].Name,
				Description: request.ContestList[i].Description,
			}
			contestList = append(contestList, &contest)
		}
		for i := 0; i < len(request.CertificateList); i++ {
			var certificate = model.ContestCertificate{
				Id:          resume1.GlobalEngine.Options.Id.GenId(),
				ResumeId:    resumeId,
				Name:        request.CertificateList[i].Name,
				Description: request.CertificateList[i].Description,
			}
			contestList = append(contestList, &certificate)
		}
		var contestListValue = []interface{}{contestList}
		value = append(value, contestListValue)
		length++
	}

	languageList := make([]*model.Language, 0, len(request.LanguageList))
	if len(request.LanguageList) != 0 {
		for i := 0; i < len(request.LanguageList); i++ {
			var language = model.Language{
				Id:               resume1.GlobalEngine.Options.Id.GenId(),
				ResumeId:         resumeId,
				LanguageName:     request.LanguageList[i].Language,
				ProficiencyLevel: request.LanguageList[i].ProficiencyLevel,
			}
			languageList = append(languageList, &language)
		}
		var languageListValue = []interface{}{languageList}
		value = append(value, languageListValue)
		length++
	}

	socialList := make([]*model.Social, 0, len(request.SocialList))
	if len(request.SocialList) != 0 {
		for i := 0; i < len(request.SocialList); i++ {
			var social = model.Social{
				Id:             resume1.GlobalEngine.Options.Id.GenId(),
				ResumeId:       resumeId,
				SocialPlatform: request.SocialList[i].SocialPlatform,
				PlatformUrl:    request.SocialList[i].PlatformUrl,
			}
			socialList = append(socialList, &social)
		}
		var socialListValue = []interface{}{socialList}
		value = append(value, socialListValue)
		length++
	}

	fs := make([]string, 0, length)
	for i := 0; i < length; i++ {
		fs = append(fs, pkg.DbFunctionCreate)
	}

	if err := resume1.GlobalEngine.Options.Orm.Transaction(fs, value); err != nil {
		klog.Error(errCreateResume, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

func GetResumeById(request *resume.GetResumeByIdRPCRequest) (int32, string, *resume.ResumeInfo, error) {
	var resumeBasicInfo model.BasicInfo
	isExist, err := resume1.GlobalEngine.Options.Orm.Query(
		1,
		-1,
		&resumeBasicInfo,
		"created_at asc",
		[]string{"resume_id", "user_id", "name", "phone", "resume_url", "email", "self_evaluation", "birthday"},
		"user_id = ?",
		request.UserId)
	if err != nil || isExist == false {
		klog.Error(errGetResumeById, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	if resumeBasicInfo.ResumeId == 0 {
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil, nil
	}

	chEducationalExperienceList := make(chan []model.EducationalExperience, 0)
	chInternshipExperienceList := make(chan []model.InternshipWorkExperience, 0)
	chWorkExperienceList := make(chan []model.InternshipWorkExperience, 0)
	chProjectExperienceList := make(chan []model.ProjectExperience, 0)
	chContestList := make(chan []model.ContestCertificate, 0)
	chCertificateList := make(chan []model.ContestCertificate, 0)
	chLanguageList := make(chan []model.Language, 0)
	chSocialList := make(chan []model.Social, 0)
	chTalentPortrait := make(chan []model.TalentPortrait, 0)
	errChannel := make(chan error, 0)

	getResumeInfoWithoutBasic([]int64{resumeBasicInfo.ResumeId}, []int64{request.UserId}, chEducationalExperienceList, chInternshipExperienceList,
		chWorkExperienceList, chProjectExperienceList, chContestList, chCertificateList, chLanguageList, chSocialList,
		chTalentPortrait, errChannel)

	count := 0
	var educationalExperienceList []model.EducationalExperience
	var internshipExperienceList []model.InternshipWorkExperience
	var workExperienceList []model.InternshipWorkExperience
	var projectExperienceList []model.ProjectExperience
	var contestList []model.ContestCertificate
	var certificateList []model.ContestCertificate
	var languageList []model.Language
	var socialList []model.Social
	var talent []model.TalentPortrait

	for {
		flag := 0
		select {
		case err = <-errChannel:
			flag = 1
			break
		case <-time.After(time.Second * 5):
			err = errors.New(errTimeOut)
			flag = 1
			break
		case educationalExperienceList = <-chEducationalExperienceList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case internshipExperienceList = <-chInternshipExperienceList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case workExperienceList = <-chWorkExperienceList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case projectExperienceList = <-chProjectExperienceList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case contestList = <-chContestList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case certificateList = <-chCertificateList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case languageList = <-chLanguageList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case socialList = <-chSocialList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case talent = <-chTalentPortrait:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		}
		if flag == 1 {
			break
		}
	}

	if err != nil {
		klog.Error(errGetResumeById, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	var basicInfo = base.BasicInfoWithResumeId{
		Name:          resumeBasicInfo.Name,
		Phone:         resumeBasicInfo.Phone,
		Email:         resumeBasicInfo.Email,
		SelfEvaluaton: resumeBasicInfo.SelfEvaluation,
		Birthday:      resumeBasicInfo.Birthday.UnixMilli(),
		UserId:        resumeBasicInfo.UserId,
		ResumeUrl:     resumeBasicInfo.ResumeUrl,
		ResumeId:      resumeBasicInfo.ResumeId,
	}

	var resultTalent = resume.TalentPortrait{
		Age:              int16(talent[0].Age),
		MaxEducation:     talent[0].MaxEducation,
		GraduatedSchool:  talent[0].GraduatedSchool,
		SchoolLevel:      talent[0].SchoolLevel,
		WorkingSeniority: int16(talent[0].WorkingSeniority),
	}

	EducationalExperienceList := make([]*base.EducationalExperienceList, 0, len(educationalExperienceList))
	for i := 0; i < len(educationalExperienceList); i++ {
		var education = base.EducationalExperienceList{
			School:     educationalExperienceList[i].School,
			Education:  educationalExperienceList[i].Education,
			Speciality: educationalExperienceList[i].Speciality,
			StartTime:  educationalExperienceList[i].StartTime.UnixMilli(),
			EndTime:    educationalExperienceList[i].EndTime.UnixMilli(),
			Rank:       educationalExperienceList[i].Ranking,
		}
		EducationalExperienceList = append(EducationalExperienceList, &education)
	}

	InternshipExperienceList := make([]*base.InternshipExperienceList, 0, len(internshipExperienceList))
	for i := 0; i < len(internshipExperienceList); i++ {
		var internship = base.InternshipExperienceList{
			Company:     internshipExperienceList[i].Company,
			Position:    internshipExperienceList[i].Position,
			Description: internshipExperienceList[i].Description,
			StartTime:   internshipExperienceList[i].StartTime.UnixMilli(),
			EndTime:     internshipExperienceList[i].EndTime.UnixMilli(),
		}
		InternshipExperienceList = append(InternshipExperienceList, &internship)
	}

	WorkExperienceList := make([]*base.WorkExperienceList, 0, len(workExperienceList))
	for i := 0; i < len(workExperienceList); i++ {
		var workExperience = base.WorkExperienceList{
			Company:     workExperienceList[i].Company,
			Position:    workExperienceList[i].Position,
			Description: workExperienceList[i].Description,
			StartTime:   workExperienceList[i].StartTime.UnixMilli(),
			EndTime:     workExperienceList[i].EndTime.UnixMilli(),
		}
		WorkExperienceList = append(WorkExperienceList, &workExperience)
	}

	ProjectExperienceList := make([]*base.ProjectExperienceList, 0, len(projectExperienceList))
	for i := 0; i < len(projectExperienceList); i++ {
		var project = base.ProjectExperienceList{
			ProjectName:        projectExperienceList[i].ProjectName,
			ProjectRole:        projectExperienceList[i].ProjectRole,
			StartTime:          projectExperienceList[i].StartTime.UnixMilli(),
			EndTime:            projectExperienceList[i].EndTime.UnixMilli(),
			ProjectDescription: projectExperienceList[i].ProjectDescription,
			ProjectUrl:         projectExperienceList[i].ProjectUrl,
		}
		ProjectExperienceList = append(ProjectExperienceList, &project)
	}

	ContestList := make([]*base.ContestList, 0, len(contestList))
	for i := 0; i < len(contestList); i++ {
		var contest = base.ContestList{
			Name:        contestList[i].Name,
			Description: contestList[i].Description,
		}
		ContestList = append(ContestList, &contest)
	}

	CertificateList := make([]*base.CertificateList, 0, len(certificateList))
	for i := 0; i < len(certificateList); i++ {
		var certificate = base.CertificateList{
			Name:        certificateList[i].Name,
			Description: certificateList[i].Description,
		}
		CertificateList = append(CertificateList, &certificate)
	}

	LanguageList := make([]*base.LanguageList, 0, len(languageList))
	for i := 0; i < len(languageList); i++ {
		var language = base.LanguageList{
			Language:         languageList[i].LanguageName,
			ProficiencyLevel: languageList[i].ProficiencyLevel,
		}
		LanguageList = append(LanguageList, &language)
	}

	SocialList := make([]*base.SocialList, 0, len(socialList))
	for i := 0; i < len(socialList); i++ {
		var social = base.SocialList{
			SocialPlatform: socialList[i].SocialPlatform,
			PlatformUrl:    socialList[i].PlatformUrl,
		}
		SocialList = append(SocialList, &social)
	}

	var resultData = resume.ResumeInfo{
		BasicInformation:          &basicInfo,
		EducationalExperienceList: EducationalExperienceList,
		InternshipExperienceList:  InternshipExperienceList,
		WorkExperienceList:        WorkExperienceList,
		ProjectExperienceList:     ProjectExperienceList,
		ContestList:               ContestList,
		CertificateList:           CertificateList,
		LanguageList:              LanguageList,
		SocialList:                SocialList,
		TalentPortrait:            &resultTalent,
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), &resultData, nil
}

func GetResumeByPost(request *resume.GetResumeByPostRPCRequest) (int32, string, []*resume.ResumeInfo, error) {
	response, err := resume1.GlobalEngine.Options.PostClient.Client.GetUserIdByPostId(context.Background(), &post.GetUserIdByPostIdRPCRequest{
		IsTalentPool: request.IsTalentPool,
		PostId:       request.PostId,
		Limit:        request.Limit,
		Offset:       request.Offset,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil
	}
	if response.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return response.Code, response.Message, nil, nil
	}

	if len(response.UserIdList) == 0 {
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil, nil
	}

	resumeBasicInfo := make([]model.BasicInfo, 0, len(response.UserIdList))
	isExist, err := resume1.GlobalEngine.Options.Orm.Query(
		-1,
		-1,
		&resumeBasicInfo,
		"",
		[]string{"resume_id", "user_id", "name", "phone", "resume_url", "email", "self_evaluation", "birthday"},
		"user_id in ?",
		response.UserIdList)
	if err != nil || isExist == false {
		klog.Error(errGetResumeByPost, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	resumeIdList := make([]int64, 0, len(resumeBasicInfo))
	for i := 0; i < len(resumeBasicInfo); i++ {
		resumeIdList = append(resumeIdList, resumeBasicInfo[i].ResumeId)
	}

	chEducationalExperienceList := make(chan []model.EducationalExperience, 0)
	chInternshipExperienceList := make(chan []model.InternshipWorkExperience, 0)
	chWorkExperienceList := make(chan []model.InternshipWorkExperience, 0)
	chProjectExperienceList := make(chan []model.ProjectExperience, 0)
	chContestList := make(chan []model.ContestCertificate, 0)
	chCertificateList := make(chan []model.ContestCertificate, 0)
	chLanguageList := make(chan []model.Language, 0)
	chSocialList := make(chan []model.Social, 0)
	chTalentPortrait := make(chan []model.TalentPortrait, 0)
	errChannel := make(chan error, 0)

	getResumeInfoWithoutBasic(resumeIdList, response.UserIdList, chEducationalExperienceList, chInternshipExperienceList,
		chWorkExperienceList, chProjectExperienceList, chContestList, chCertificateList, chLanguageList, chSocialList,
		chTalentPortrait, errChannel)

	count := 0
	var educationalExperienceList []model.EducationalExperience
	var internshipExperienceList []model.InternshipWorkExperience
	var workExperienceList []model.InternshipWorkExperience
	var projectExperienceList []model.ProjectExperience
	var contestList []model.ContestCertificate
	var certificateList []model.ContestCertificate
	var languageList []model.Language
	var socialList []model.Social
	var talent []model.TalentPortrait

	for {
		flag := 0
		select {
		case err = <-errChannel:
			flag = 1
			break
		case <-time.After(time.Second * 5):
			err = errors.New(errTimeOut)
			flag = 1
			break
		case educationalExperienceList = <-chEducationalExperienceList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case internshipExperienceList = <-chInternshipExperienceList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case workExperienceList = <-chWorkExperienceList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case projectExperienceList = <-chProjectExperienceList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case contestList = <-chContestList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case certificateList = <-chCertificateList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case languageList = <-chLanguageList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case socialList = <-chSocialList:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		case talent = <-chTalentPortrait:
			count++
			if count == breakLength {
				flag = 1
				break
			}
		}
		if flag == 1 {
			break
		}
	}

	if err != nil {
		klog.Error(errGetResumeByPost, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	resultDatas := make([]*resume.ResumeInfo, 0, len(response.UserIdList))

	for i := 0; i < len(resumeBasicInfo); i++ {
		var basicInfo = base.BasicInfoWithResumeId{
			Name:          resumeBasicInfo[i].Name,
			Phone:         resumeBasicInfo[i].Phone,
			Email:         resumeBasicInfo[i].Email,
			SelfEvaluaton: resumeBasicInfo[i].SelfEvaluation,
			Birthday:      resumeBasicInfo[i].Birthday.UnixMilli(),
			UserId:        resumeBasicInfo[i].UserId,
			ResumeUrl:     resumeBasicInfo[i].ResumeUrl,
			ResumeId:      resumeBasicInfo[i].ResumeId,
		}

		var resultTalent = resume.TalentPortrait{
			Age:              int16(talent[i].Age),
			MaxEducation:     talent[i].MaxEducation,
			GraduatedSchool:  talent[i].GraduatedSchool,
			SchoolLevel:      talent[i].SchoolLevel,
			WorkingSeniority: int16(talent[i].WorkingSeniority),
		}

		EducationalExperienceList := make([]*base.EducationalExperienceList, 0, len(educationalExperienceList))
		for j := 0; j < len(educationalExperienceList); j++ {
			if educationalExperienceList[j].ResumeId == resumeBasicInfo[i].ResumeId {
				var education = base.EducationalExperienceList{
					School:     educationalExperienceList[j].School,
					Education:  educationalExperienceList[j].Education,
					Speciality: educationalExperienceList[j].Speciality,
					StartTime:  educationalExperienceList[j].StartTime.UnixMilli(),
					EndTime:    educationalExperienceList[j].EndTime.UnixMilli(),
					Rank:       educationalExperienceList[j].Ranking,
				}
				EducationalExperienceList = append(EducationalExperienceList, &education)
			}
		}

		InternshipExperienceList := make([]*base.InternshipExperienceList, 0, len(internshipExperienceList))
		for j := 0; j < len(internshipExperienceList); j++ {
			if internshipExperienceList[j].ResumeId == resumeBasicInfo[i].ResumeId {
				var internship = base.InternshipExperienceList{
					Company:     internshipExperienceList[j].Company,
					Position:    internshipExperienceList[j].Position,
					Description: internshipExperienceList[j].Description,
					StartTime:   internshipExperienceList[j].StartTime.UnixMilli(),
					EndTime:     internshipExperienceList[j].EndTime.UnixMilli(),
				}
				InternshipExperienceList = append(InternshipExperienceList, &internship)
			}
		}

		WorkExperienceList := make([]*base.WorkExperienceList, 0, len(workExperienceList))
		for j := 0; j < len(workExperienceList); j++ {
			if workExperienceList[j].ResumeId == resumeBasicInfo[i].ResumeId {
				var workExperience = base.WorkExperienceList{
					Company:     workExperienceList[j].Company,
					Position:    workExperienceList[j].Position,
					Description: workExperienceList[j].Description,
					StartTime:   workExperienceList[j].StartTime.UnixMilli(),
					EndTime:     workExperienceList[j].EndTime.UnixMilli(),
				}
				WorkExperienceList = append(WorkExperienceList, &workExperience)
			}
		}

		ProjectExperienceList := make([]*base.ProjectExperienceList, 0, len(projectExperienceList))
		for j := 0; j < len(projectExperienceList); j++ {
			if projectExperienceList[j].ResumeId == resumeBasicInfo[i].ResumeId {
				var project = base.ProjectExperienceList{
					ProjectName:        projectExperienceList[j].ProjectName,
					ProjectRole:        projectExperienceList[j].ProjectRole,
					StartTime:          projectExperienceList[j].StartTime.UnixMilli(),
					EndTime:            projectExperienceList[j].EndTime.UnixMilli(),
					ProjectDescription: projectExperienceList[j].ProjectDescription,
					ProjectUrl:         projectExperienceList[j].ProjectUrl,
				}
				ProjectExperienceList = append(ProjectExperienceList, &project)
			}
		}

		ContestList := make([]*base.ContestList, 0, len(contestList))
		for j := 0; j < len(contestList); j++ {
			if contestList[j].ResumeId == resumeBasicInfo[i].ResumeId {
				var contest = base.ContestList{
					Name:        contestList[j].Name,
					Description: contestList[j].Description,
				}
				ContestList = append(ContestList, &contest)
			}
		}

		CertificateList := make([]*base.CertificateList, 0, len(certificateList))
		for j := 0; j < len(certificateList); j++ {
			if certificateList[j].ResumeId == resumeBasicInfo[i].ResumeId {
				var certificate = base.CertificateList{
					Name:        certificateList[j].Name,
					Description: certificateList[j].Description,
				}
				CertificateList = append(CertificateList, &certificate)
			}
		}

		LanguageList := make([]*base.LanguageList, 0, len(languageList))
		for j := 0; j < len(languageList); j++ {
			if languageList[j].ResumeId == resumeBasicInfo[i].ResumeId {
				var language = base.LanguageList{
					Language:         languageList[j].LanguageName,
					ProficiencyLevel: languageList[j].ProficiencyLevel,
				}
				LanguageList = append(LanguageList, &language)
			}
		}

		SocialList := make([]*base.SocialList, 0, len(socialList))
		for j := 0; j < len(socialList); j++ {
			if socialList[j].ResumeId == resumeBasicInfo[i].ResumeId {
				var social = base.SocialList{
					SocialPlatform: socialList[j].SocialPlatform,
					PlatformUrl:    socialList[j].PlatformUrl,
				}
				SocialList = append(SocialList, &social)
			}
		}

		var resultData = resume.ResumeInfo{
			BasicInformation:          &basicInfo,
			EducationalExperienceList: EducationalExperienceList,
			InternshipExperienceList:  InternshipExperienceList,
			WorkExperienceList:        WorkExperienceList,
			ProjectExperienceList:     ProjectExperienceList,
			ContestList:               ContestList,
			CertificateList:           CertificateList,
			LanguageList:              LanguageList,
			SocialList:                SocialList,
			TalentPortrait:            &resultTalent,
		}

		resultDatas = append(resultDatas, &resultData)
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil

}

func getResumeInfoWithoutBasic(resumeId, userId []int64,
	chEducationalExperienceList chan []model.EducationalExperience,
	chInternshipExperienceList chan []model.InternshipWorkExperience,
	chWorkExperienceList chan []model.InternshipWorkExperience,
	chProjectExperienceList chan []model.ProjectExperience,
	chContestList chan []model.ContestCertificate,
	chCertificateList chan []model.ContestCertificate,
	chLanguageList chan []model.Language,
	chSocialList chan []model.Social,
	chTalentPortrait chan []model.TalentPortrait,
	errChannel chan error) {
	go func() {
		talent := make([]model.TalentPortrait, 0, len(userId))
		isExist, err := resume1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&talent,
			"",
			[]string{"age", "max_education", "graduated_school", "school_level", "working_seniority"},
			"user_id in ?",
			userId)
		if err != nil || isExist == false {
			klog.Error(errGetResumeById, err)
			errChannel <- err
			return
		}

		chTalentPortrait <- talent
	}()

	go func() {
		educationalExperienceList := make([]model.EducationalExperience, 0, maxLength)
		_, err := resume1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&educationalExperienceList,
			"",
			[]string{"resume_id", "school", "education", "speciality", "ranking", "start_time", "end_time"},
			"resume_id in ?",
			resumeId)
		if err != nil {
			klog.Error(errGetResumeById, err)
			errChannel <- err
			return
		}

		chEducationalExperienceList <- educationalExperienceList
	}()

	go func() {
		internshipExperienceList := make([]model.InternshipWorkExperience, 0, maxLength)
		_, err := resume1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&internshipExperienceList,
			"",
			[]string{"resume_id", "company", "position", "description", "start_time", "end_time"},
			"resume_id in ? and is_internship = ?",
			resumeId, 0)
		if err != nil {
			klog.Error(errGetResumeById, err)
			errChannel <- err
			return
		}

		chInternshipExperienceList <- internshipExperienceList
	}()

	go func() {
		workExperienceList := make([]model.InternshipWorkExperience, 0, maxLength)
		_, err := resume1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&workExperienceList,
			"",
			[]string{"resume_id", "company", "position", "description", "start_time", "end_time"},
			"resume_id in ? and is_internship = ?",
			resumeId, 1)
		if err != nil {
			klog.Error(errGetResumeById, err)
			errChannel <- err
			return
		}

		chWorkExperienceList <- workExperienceList
	}()

	go func() {
		projectExperienceList := make([]model.ProjectExperience, 0, maxLength)
		_, err := resume1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&projectExperienceList,
			"",
			[]string{"resume_id", "project_name", "project_role", "project_description", "project_url", "start_time", "end_time"},
			"resume_id in ?",
			resumeId)
		if err != nil {
			klog.Error(errGetResumeById, err)
			errChannel <- err
			return
		}

		chProjectExperienceList <- projectExperienceList
	}()

	go func() {
		contestList := make([]model.ContestCertificate, 0, maxLength)
		_, err := resume1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&contestList,
			"",
			[]string{"resume_id", "name", "description"},
			"resume_id in ? and is_contest = ?",
			resumeId, 0)
		if err != nil {
			klog.Error(errGetResumeById, err)
			errChannel <- err
			return
		}

		chContestList <- contestList
	}()

	go func() {
		certificateList := make([]model.ContestCertificate, 0, maxLength)
		_, err := resume1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&certificateList,
			"",
			[]string{"resume_id", "name", "description"},
			"resume_id in ? and is_contest = ?",
			resumeId, 1)
		if err != nil {
			klog.Error(errGetResumeById, err)
			errChannel <- err
			return
		}

		chCertificateList <- certificateList
	}()

	go func() {
		languageList := make([]model.Language, 0, maxLength)
		_, err := resume1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&languageList,
			"",
			[]string{"resume_id", "language_name", "proficiency_level"},
			"resume_id in ?",
			resumeId)
		if err != nil {
			klog.Error(errGetResumeById, err)
			errChannel <- err
			return
		}

		chLanguageList <- languageList
	}()

	go func() {
		socialList := make([]model.Social, 0, maxLength)
		_, err := resume1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&socialList,
			"",
			[]string{"resume_id", "social_platform", "platform_url"},
			"resume_id in ?",
			resumeId)
		if err != nil {
			klog.Error(errGetResumeById, err)
			errChannel <- err
			return
		}

		chSocialList <- socialList
	}()
}
