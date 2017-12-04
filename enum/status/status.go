package status

//go:generate stringer -type=Status
type Status int

const (
	OperationWasSuccessful                     Status = 0
	TheUserIDProvidedIsAbsentOrIncorrect       Status = 247
	TheProvidedUserIDAndOrOauthCredsDoNotMatch Status = 250
	TokenIsInvalidOrDoesntExist                Status = 283
	NoSuchSubscription                         Status = 286
	TheCallbackURLIsEitherAbsentOrIncorrect    Status = 283
	NoSuchSubscriptionCouldBeDeleted           Status = 294
	CommentAbsentOrIncorrect                   Status = 304
	TooManyNotificationsSet                    Status = 305
	UserIsDeactiviated                         Status = 328
	SignatureIsInvalid                         Status = 342
	WrongNotificationCallbackURL               Status = 343
	TooManyRequets                             Status = 601
	WrongActionOrWrongWebservice               Status = 2554
	UnknonwError                               Status = 2555
	ServiceNotDefined                          Status = 2556
)
