package types

// PassportElementError Represents an error in the Telegram Passport element which was submitted that should be resolved by the user
// It should be one of:
//  - PassportElementErrorDataField
//  - PassportElementErrorFrontSide
//  - PassportElementErrorReverseSide
//  - PassportElementErrorSelfie
//  - PassportElementErrorFile
//  - PassportElementErrorFiles
//  - PassportElementErrorTranslationFile
//  - PassportElementErrorTranslationFiles
//  - PassportElementErrorUnspecified
type PassportElementError interface{}
