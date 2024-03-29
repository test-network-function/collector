package util

const InvalidRequestErr = "Invalid request."
const RequestContentTypeErr = "Error found while Parsing multipart form: %s"
const WritingResponseErr = "Error found while writing response: %s"
const FormFileErr = "Error found while forming file: %s"
const UnmarshalErr = "Error found while trying to unmarshal claim file: %s"
const MarshalErr = "Error found while marshaling claim file: %s"
const MalformedClaimFileErr = "Malformed claim file: %s"
const ClaimFieldMissingErr = "claim field is missing"
const VersionsFieldMissingErr = "versions field is missing"
const OcpFieldMissingErr = "ocp subfield of versions field is missing"
const TestTestIDMissingErr = "testID subfield of %s test is missing"
const TestStateMissingErr = "state subfield of %s test is missing"
const TestIDSuiteMissingErr = "suite subfield of %s's testID field is missing"
const TestIDIDMissingErr = "id subfield of %s's testID field is missing"
const ResultsFieldMissingErr = "results field is missing"
const ExecutedByMissingErr = "executed by value is missing"
const InvalidPasswordErr = "wrong password"
const InvalidUsernameErr = "partner is not found"
const RollbackErr = "Error found while Rollbacking transaction: %s"
const ExecQueryErr = "Error found while executing a mysql query: %s"
const ScanDBFieldErr = "Error found while scanning db field: %s"
const BeginTxErr = "Error found while beginning transaction: %s"
const CommitTxErr = "Error found while committing transaction: %s"
const AuthError = "Error found while authenticating partner's password: %s"
const ClaimFileError = "Error while storing the claim file in the database: %s"
const EncodingPasswordError = "Failed encoding password" // #nosec
const ServerIsUpMsg = "Server is up"
const SuccessUploadingFileMSG = "File was uploaded successfully!"
const ServerReadTimeOutEnvVarErr = "SERVER_READ_TIMEOUT environment variable must be set"
const ServerWriteTimeOutEnvVarErr = "SERVER_WRITE_TIMEOUT environment variable must be set"
const ServerAddrEnvVarErr = "SERVER_ADDR environment variable must be set"
const ServerEnvVarsError = "Error found while extracting environment variables related to the server: %s"
const PartnerOrPasswordArgsMissingErr = "partner name or Password are not given"
const EmptyClaimFileErr = "claim file is empty"
const PostRequestIsNotValidErr = "Post request isn't valid: %s"
const FailedToUploadFileToS3Err = "Error found while trying to upload claim file to bucket: %s"
const FailedToDeleteFileFromS3Err = "Error found while trying to delete file from s3: %s"

const ClaimTag = "claim"
const VersionsTag = "versions"
const ResultsTag = "results"
const ClaimFileInputName = "claimFile"
const ExecutedByInputName = "executed_by"
const PartnerNameInputName = "partner_name"
const DedcodedPasswordInputName = "decoded_password"

const UseCollectorSQLCmd = `USE cnf; `
const InsertToClaimSQLCmd = `INSERT INTO claim
								(cnf_version, executed_by, upload_time, partner_name, s3_file_key)
								VALUES (?, ?, ?, ?, ?);`
const InsertToClaimResSQLCmd = `INSERT INTO claim_result
							(claim_id, suite_name, test_id, test_status)
							VALUES (?, ?, ?, ?);`
const ExtractLastClaimID = `SELECT id FROM cnf.claim ORDER BY id DESC LIMIT 1;`
const ExtractPartnerAndPasswordCmd = `SELECT encoded_password FROM cnf.authenticator WHERE partner_name = ?`
const InsertPartnerToAuthSQLCmd = `INSERT INTO cnf.authenticator (partner_name, encoded_password) VALUES (?, ?)`
const ParseLowerBound = 10
const ParseUpperBound = 20
const FileStoredIntoClaimTableSuccessfully = `Claim is stored into table successfully.`
const FileStoredIntoClaimResultTableSuccessfully = `Claim is stored into table CLAIM_RESULT successfully.`
const FileUploadedSuccessfullyToBucket = `Claim file is successfully uploaded to bucket %s`
const FileHasBeenDeletedFromBucket = `Claim file has been deleted from bucket %s`

const SelectAllFromClaimByPartner = "SELECT * FROM cnf.claim WHERE partner_name = ?"
const SelectAllFromClaim = "SELECT * FROM cnf.claim"
const SelectAllClaimIDsByPartner = "SELECT id FROM cnf.claim WHERE partner_name = ?"
const SelectAllFromClaimResultByClaimIDs = "SELECT * FROM cnf.claim_result WHERE claim_id IN (%s)"
const SelectAllFromClaimResult = "SELECT * FROM cnf.claim_result"
const AdminUserName = "admin"
