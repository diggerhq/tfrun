// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                                = new(Query)
	AccountDeleteToken               *accountDeleteToken
	Chat                             *chat
	Customer                         *customer
	DiggerBatch                      *diggerBatch
	DiggerJob                        *diggerJob
	DiggerJobSummary                 *diggerJobSummary
	DiggerLock                       *diggerLock
	DiggerRun                        *diggerRun
	DiggerRunStage                   *diggerRunStage
	GithubApp                        *githubApp
	GithubAppInstallation            *githubAppInstallation
	GithubAppInstallationLink        *githubAppInstallationLink
	InternalBlogAuthorPost           *internalBlogAuthorPost
	InternalBlogAuthorProfile        *internalBlogAuthorProfile
	InternalBlogPost                 *internalBlogPost
	InternalBlogPostTag              *internalBlogPostTag
	InternalBlogPostTagsRelationship *internalBlogPostTagsRelationship
	InternalChangelog                *internalChangelog
	InternalFeedbackComment          *internalFeedbackComment
	InternalFeedbackThread           *internalFeedbackThread
	Organization                     *organization
	OrganizationCredit               *organizationCredit
	OrganizationJoinInvitation       *organizationJoinInvitation
	OrganizationMember               *organizationMember
	OrganizationsPrivateInfo         *organizationsPrivateInfo
	Price                            *price
	Product                          *product
	Project                          *project
	ProjectComment                   *projectComment
	Repo                             *repo
	Subscription                     *subscription
	UserAPIKey                       *userAPIKey
	UserNotification                 *userNotification
	UserOnboarding                   *userOnboarding
	UserPrivateInfo                  *userPrivateInfo
	UserProfile                      *userProfile
	UserRole                         *userRole
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AccountDeleteToken = &Q.AccountDeleteToken
	Chat = &Q.Chat
	Customer = &Q.Customer
	DiggerBatch = &Q.DiggerBatch
	DiggerJob = &Q.DiggerJob
	DiggerJobSummary = &Q.DiggerJobSummary
	DiggerLock = &Q.DiggerLock
	DiggerRun = &Q.DiggerRun
	DiggerRunStage = &Q.DiggerRunStage
	GithubApp = &Q.GithubApp
	GithubAppInstallation = &Q.GithubAppInstallation
	GithubAppInstallationLink = &Q.GithubAppInstallationLink
	InternalBlogAuthorPost = &Q.InternalBlogAuthorPost
	InternalBlogAuthorProfile = &Q.InternalBlogAuthorProfile
	InternalBlogPost = &Q.InternalBlogPost
	InternalBlogPostTag = &Q.InternalBlogPostTag
	InternalBlogPostTagsRelationship = &Q.InternalBlogPostTagsRelationship
	InternalChangelog = &Q.InternalChangelog
	InternalFeedbackComment = &Q.InternalFeedbackComment
	InternalFeedbackThread = &Q.InternalFeedbackThread
	Organization = &Q.Organization
	OrganizationCredit = &Q.OrganizationCredit
	OrganizationJoinInvitation = &Q.OrganizationJoinInvitation
	OrganizationMember = &Q.OrganizationMember
	OrganizationsPrivateInfo = &Q.OrganizationsPrivateInfo
	Price = &Q.Price
	Product = &Q.Product
	Project = &Q.Project
	ProjectComment = &Q.ProjectComment
	Repo = &Q.Repo
	Subscription = &Q.Subscription
	UserAPIKey = &Q.UserAPIKey
	UserNotification = &Q.UserNotification
	UserOnboarding = &Q.UserOnboarding
	UserPrivateInfo = &Q.UserPrivateInfo
	UserProfile = &Q.UserProfile
	UserRole = &Q.UserRole
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                               db,
		AccountDeleteToken:               newAccountDeleteToken(db, opts...),
		Chat:                             newChat(db, opts...),
		Customer:                         newCustomer(db, opts...),
		DiggerBatch:                      newDiggerBatch(db, opts...),
		DiggerJob:                        newDiggerJob(db, opts...),
		DiggerJobSummary:                 newDiggerJobSummary(db, opts...),
		DiggerLock:                       newDiggerLock(db, opts...),
		DiggerRun:                        newDiggerRun(db, opts...),
		DiggerRunStage:                   newDiggerRunStage(db, opts...),
		GithubApp:                        newGithubApp(db, opts...),
		GithubAppInstallation:            newGithubAppInstallation(db, opts...),
		GithubAppInstallationLink:        newGithubAppInstallationLink(db, opts...),
		InternalBlogAuthorPost:           newInternalBlogAuthorPost(db, opts...),
		InternalBlogAuthorProfile:        newInternalBlogAuthorProfile(db, opts...),
		InternalBlogPost:                 newInternalBlogPost(db, opts...),
		InternalBlogPostTag:              newInternalBlogPostTag(db, opts...),
		InternalBlogPostTagsRelationship: newInternalBlogPostTagsRelationship(db, opts...),
		InternalChangelog:                newInternalChangelog(db, opts...),
		InternalFeedbackComment:          newInternalFeedbackComment(db, opts...),
		InternalFeedbackThread:           newInternalFeedbackThread(db, opts...),
		Organization:                     newOrganization(db, opts...),
		OrganizationCredit:               newOrganizationCredit(db, opts...),
		OrganizationJoinInvitation:       newOrganizationJoinInvitation(db, opts...),
		OrganizationMember:               newOrganizationMember(db, opts...),
		OrganizationsPrivateInfo:         newOrganizationsPrivateInfo(db, opts...),
		Price:                            newPrice(db, opts...),
		Product:                          newProduct(db, opts...),
		Project:                          newProject(db, opts...),
		ProjectComment:                   newProjectComment(db, opts...),
		Repo:                             newRepo(db, opts...),
		Subscription:                     newSubscription(db, opts...),
		UserAPIKey:                       newUserAPIKey(db, opts...),
		UserNotification:                 newUserNotification(db, opts...),
		UserOnboarding:                   newUserOnboarding(db, opts...),
		UserPrivateInfo:                  newUserPrivateInfo(db, opts...),
		UserProfile:                      newUserProfile(db, opts...),
		UserRole:                         newUserRole(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AccountDeleteToken               accountDeleteToken
	Chat                             chat
	Customer                         customer
	DiggerBatch                      diggerBatch
	DiggerJob                        diggerJob
	DiggerJobSummary                 diggerJobSummary
	DiggerLock                       diggerLock
	DiggerRun                        diggerRun
	DiggerRunStage                   diggerRunStage
	GithubApp                        githubApp
	GithubAppInstallation            githubAppInstallation
	GithubAppInstallationLink        githubAppInstallationLink
	InternalBlogAuthorPost           internalBlogAuthorPost
	InternalBlogAuthorProfile        internalBlogAuthorProfile
	InternalBlogPost                 internalBlogPost
	InternalBlogPostTag              internalBlogPostTag
	InternalBlogPostTagsRelationship internalBlogPostTagsRelationship
	InternalChangelog                internalChangelog
	InternalFeedbackComment          internalFeedbackComment
	InternalFeedbackThread           internalFeedbackThread
	Organization                     organization
	OrganizationCredit               organizationCredit
	OrganizationJoinInvitation       organizationJoinInvitation
	OrganizationMember               organizationMember
	OrganizationsPrivateInfo         organizationsPrivateInfo
	Price                            price
	Product                          product
	Project                          project
	ProjectComment                   projectComment
	Repo                             repo
	Subscription                     subscription
	UserAPIKey                       userAPIKey
	UserNotification                 userNotification
	UserOnboarding                   userOnboarding
	UserPrivateInfo                  userPrivateInfo
	UserProfile                      userProfile
	UserRole                         userRole
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                               db,
		AccountDeleteToken:               q.AccountDeleteToken.clone(db),
		Chat:                             q.Chat.clone(db),
		Customer:                         q.Customer.clone(db),
		DiggerBatch:                      q.DiggerBatch.clone(db),
		DiggerJob:                        q.DiggerJob.clone(db),
		DiggerJobSummary:                 q.DiggerJobSummary.clone(db),
		DiggerLock:                       q.DiggerLock.clone(db),
		DiggerRun:                        q.DiggerRun.clone(db),
		DiggerRunStage:                   q.DiggerRunStage.clone(db),
		GithubApp:                        q.GithubApp.clone(db),
		GithubAppInstallation:            q.GithubAppInstallation.clone(db),
		GithubAppInstallationLink:        q.GithubAppInstallationLink.clone(db),
		InternalBlogAuthorPost:           q.InternalBlogAuthorPost.clone(db),
		InternalBlogAuthorProfile:        q.InternalBlogAuthorProfile.clone(db),
		InternalBlogPost:                 q.InternalBlogPost.clone(db),
		InternalBlogPostTag:              q.InternalBlogPostTag.clone(db),
		InternalBlogPostTagsRelationship: q.InternalBlogPostTagsRelationship.clone(db),
		InternalChangelog:                q.InternalChangelog.clone(db),
		InternalFeedbackComment:          q.InternalFeedbackComment.clone(db),
		InternalFeedbackThread:           q.InternalFeedbackThread.clone(db),
		Organization:                     q.Organization.clone(db),
		OrganizationCredit:               q.OrganizationCredit.clone(db),
		OrganizationJoinInvitation:       q.OrganizationJoinInvitation.clone(db),
		OrganizationMember:               q.OrganizationMember.clone(db),
		OrganizationsPrivateInfo:         q.OrganizationsPrivateInfo.clone(db),
		Price:                            q.Price.clone(db),
		Product:                          q.Product.clone(db),
		Project:                          q.Project.clone(db),
		ProjectComment:                   q.ProjectComment.clone(db),
		Repo:                             q.Repo.clone(db),
		Subscription:                     q.Subscription.clone(db),
		UserAPIKey:                       q.UserAPIKey.clone(db),
		UserNotification:                 q.UserNotification.clone(db),
		UserOnboarding:                   q.UserOnboarding.clone(db),
		UserPrivateInfo:                  q.UserPrivateInfo.clone(db),
		UserProfile:                      q.UserProfile.clone(db),
		UserRole:                         q.UserRole.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                               db,
		AccountDeleteToken:               q.AccountDeleteToken.replaceDB(db),
		Chat:                             q.Chat.replaceDB(db),
		Customer:                         q.Customer.replaceDB(db),
		DiggerBatch:                      q.DiggerBatch.replaceDB(db),
		DiggerJob:                        q.DiggerJob.replaceDB(db),
		DiggerJobSummary:                 q.DiggerJobSummary.replaceDB(db),
		DiggerLock:                       q.DiggerLock.replaceDB(db),
		DiggerRun:                        q.DiggerRun.replaceDB(db),
		DiggerRunStage:                   q.DiggerRunStage.replaceDB(db),
		GithubApp:                        q.GithubApp.replaceDB(db),
		GithubAppInstallation:            q.GithubAppInstallation.replaceDB(db),
		GithubAppInstallationLink:        q.GithubAppInstallationLink.replaceDB(db),
		InternalBlogAuthorPost:           q.InternalBlogAuthorPost.replaceDB(db),
		InternalBlogAuthorProfile:        q.InternalBlogAuthorProfile.replaceDB(db),
		InternalBlogPost:                 q.InternalBlogPost.replaceDB(db),
		InternalBlogPostTag:              q.InternalBlogPostTag.replaceDB(db),
		InternalBlogPostTagsRelationship: q.InternalBlogPostTagsRelationship.replaceDB(db),
		InternalChangelog:                q.InternalChangelog.replaceDB(db),
		InternalFeedbackComment:          q.InternalFeedbackComment.replaceDB(db),
		InternalFeedbackThread:           q.InternalFeedbackThread.replaceDB(db),
		Organization:                     q.Organization.replaceDB(db),
		OrganizationCredit:               q.OrganizationCredit.replaceDB(db),
		OrganizationJoinInvitation:       q.OrganizationJoinInvitation.replaceDB(db),
		OrganizationMember:               q.OrganizationMember.replaceDB(db),
		OrganizationsPrivateInfo:         q.OrganizationsPrivateInfo.replaceDB(db),
		Price:                            q.Price.replaceDB(db),
		Product:                          q.Product.replaceDB(db),
		Project:                          q.Project.replaceDB(db),
		ProjectComment:                   q.ProjectComment.replaceDB(db),
		Repo:                             q.Repo.replaceDB(db),
		Subscription:                     q.Subscription.replaceDB(db),
		UserAPIKey:                       q.UserAPIKey.replaceDB(db),
		UserNotification:                 q.UserNotification.replaceDB(db),
		UserOnboarding:                   q.UserOnboarding.replaceDB(db),
		UserPrivateInfo:                  q.UserPrivateInfo.replaceDB(db),
		UserProfile:                      q.UserProfile.replaceDB(db),
		UserRole:                         q.UserRole.replaceDB(db),
	}
}

type queryCtx struct {
	AccountDeleteToken               IAccountDeleteTokenDo
	Chat                             IChatDo
	Customer                         ICustomerDo
	DiggerBatch                      IDiggerBatchDo
	DiggerJob                        IDiggerJobDo
	DiggerJobSummary                 IDiggerJobSummaryDo
	DiggerLock                       IDiggerLockDo
	DiggerRun                        IDiggerRunDo
	DiggerRunStage                   IDiggerRunStageDo
	GithubApp                        IGithubAppDo
	GithubAppInstallation            IGithubAppInstallationDo
	GithubAppInstallationLink        IGithubAppInstallationLinkDo
	InternalBlogAuthorPost           IInternalBlogAuthorPostDo
	InternalBlogAuthorProfile        IInternalBlogAuthorProfileDo
	InternalBlogPost                 IInternalBlogPostDo
	InternalBlogPostTag              IInternalBlogPostTagDo
	InternalBlogPostTagsRelationship IInternalBlogPostTagsRelationshipDo
	InternalChangelog                IInternalChangelogDo
	InternalFeedbackComment          IInternalFeedbackCommentDo
	InternalFeedbackThread           IInternalFeedbackThreadDo
	Organization                     IOrganizationDo
	OrganizationCredit               IOrganizationCreditDo
	OrganizationJoinInvitation       IOrganizationJoinInvitationDo
	OrganizationMember               IOrganizationMemberDo
	OrganizationsPrivateInfo         IOrganizationsPrivateInfoDo
	Price                            IPriceDo
	Product                          IProductDo
	Project                          IProjectDo
	ProjectComment                   IProjectCommentDo
	Repo                             IRepoDo
	Subscription                     ISubscriptionDo
	UserAPIKey                       IUserAPIKeyDo
	UserNotification                 IUserNotificationDo
	UserOnboarding                   IUserOnboardingDo
	UserPrivateInfo                  IUserPrivateInfoDo
	UserProfile                      IUserProfileDo
	UserRole                         IUserRoleDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AccountDeleteToken:               q.AccountDeleteToken.WithContext(ctx),
		Chat:                             q.Chat.WithContext(ctx),
		Customer:                         q.Customer.WithContext(ctx),
		DiggerBatch:                      q.DiggerBatch.WithContext(ctx),
		DiggerJob:                        q.DiggerJob.WithContext(ctx),
		DiggerJobSummary:                 q.DiggerJobSummary.WithContext(ctx),
		DiggerLock:                       q.DiggerLock.WithContext(ctx),
		DiggerRun:                        q.DiggerRun.WithContext(ctx),
		DiggerRunStage:                   q.DiggerRunStage.WithContext(ctx),
		GithubApp:                        q.GithubApp.WithContext(ctx),
		GithubAppInstallation:            q.GithubAppInstallation.WithContext(ctx),
		GithubAppInstallationLink:        q.GithubAppInstallationLink.WithContext(ctx),
		InternalBlogAuthorPost:           q.InternalBlogAuthorPost.WithContext(ctx),
		InternalBlogAuthorProfile:        q.InternalBlogAuthorProfile.WithContext(ctx),
		InternalBlogPost:                 q.InternalBlogPost.WithContext(ctx),
		InternalBlogPostTag:              q.InternalBlogPostTag.WithContext(ctx),
		InternalBlogPostTagsRelationship: q.InternalBlogPostTagsRelationship.WithContext(ctx),
		InternalChangelog:                q.InternalChangelog.WithContext(ctx),
		InternalFeedbackComment:          q.InternalFeedbackComment.WithContext(ctx),
		InternalFeedbackThread:           q.InternalFeedbackThread.WithContext(ctx),
		Organization:                     q.Organization.WithContext(ctx),
		OrganizationCredit:               q.OrganizationCredit.WithContext(ctx),
		OrganizationJoinInvitation:       q.OrganizationJoinInvitation.WithContext(ctx),
		OrganizationMember:               q.OrganizationMember.WithContext(ctx),
		OrganizationsPrivateInfo:         q.OrganizationsPrivateInfo.WithContext(ctx),
		Price:                            q.Price.WithContext(ctx),
		Product:                          q.Product.WithContext(ctx),
		Project:                          q.Project.WithContext(ctx),
		ProjectComment:                   q.ProjectComment.WithContext(ctx),
		Repo:                             q.Repo.WithContext(ctx),
		Subscription:                     q.Subscription.WithContext(ctx),
		UserAPIKey:                       q.UserAPIKey.WithContext(ctx),
		UserNotification:                 q.UserNotification.WithContext(ctx),
		UserOnboarding:                   q.UserOnboarding.WithContext(ctx),
		UserPrivateInfo:                  q.UserPrivateInfo.WithContext(ctx),
		UserProfile:                      q.UserProfile.WithContext(ctx),
		UserRole:                         q.UserRole.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
