// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: GithubPullGetter)

package mocks

import (
	github "github.com/google/go-github/v64/github"
	pegomock "github.com/petergtz/pegomock/v4"
	models "github.com/runatlantis/atlantis/server/events/models"
	logging "github.com/runatlantis/atlantis/server/logging"
	"reflect"
	"time"
)

type MockGithubPullGetter struct {
	fail func(message string, callerSkip ...int)
}

func NewMockGithubPullGetter(options ...pegomock.Option) *MockGithubPullGetter {
	mock := &MockGithubPullGetter{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockGithubPullGetter) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockGithubPullGetter) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockGithubPullGetter) GetPullRequest(logger logging.SimpleLogging, repo models.Repo, pullNum int) (*github.PullRequest, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockGithubPullGetter().")
	}
	params := []pegomock.Param{logger, repo, pullNum}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetPullRequest", params, []reflect.Type{reflect.TypeOf((**github.PullRequest)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *github.PullRequest
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*github.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockGithubPullGetter) VerifyWasCalledOnce() *VerifierMockGithubPullGetter {
	return &VerifierMockGithubPullGetter{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockGithubPullGetter) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockGithubPullGetter {
	return &VerifierMockGithubPullGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockGithubPullGetter) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockGithubPullGetter {
	return &VerifierMockGithubPullGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockGithubPullGetter) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockGithubPullGetter {
	return &VerifierMockGithubPullGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockGithubPullGetter struct {
	mock                   *MockGithubPullGetter
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockGithubPullGetter) GetPullRequest(logger logging.SimpleLogging, repo models.Repo, pullNum int) *MockGithubPullGetter_GetPullRequest_OngoingVerification {
	params := []pegomock.Param{logger, repo, pullNum}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetPullRequest", params, verifier.timeout)
	return &MockGithubPullGetter_GetPullRequest_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockGithubPullGetter_GetPullRequest_OngoingVerification struct {
	mock              *MockGithubPullGetter
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockGithubPullGetter_GetPullRequest_OngoingVerification) GetCapturedArguments() (logging.SimpleLogging, models.Repo, int) {
	logger, repo, pullNum := c.GetAllCapturedArguments()
	return logger[len(logger)-1], repo[len(repo)-1], pullNum[len(pullNum)-1]
}

func (c *MockGithubPullGetter_GetPullRequest_OngoingVerification) GetAllCapturedArguments() (_param0 []logging.SimpleLogging, _param1 []models.Repo, _param2 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]logging.SimpleLogging, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(logging.SimpleLogging)
		}
		_param1 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.Repo)
		}
		_param2 = make([]int, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(int)
		}
	}
	return
}
