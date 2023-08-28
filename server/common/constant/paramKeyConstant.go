package constant

const (
	/**
	 * WeIdService related param names.
	 */
	PUBLIC_KEY = "publicKey"

	/**
	 * AuthorityIssuer related param names.
	 */
	AUTHORITY_ISSUER_NAME = "name"

	/**
	 * CptService related param names.
	 */
	CPT_JSON_SCHEMA = "cptJsonSchema"
	CPT_SIGNATURE   = "cptSignature"
	CPT             = "Cpt"

	/**
	 * CredentialService related param names.
	 */
	CPT_ID               = "cptId"
	ISSUER               = "issuer"
	CLAIM                = "claim"
	EXPIRATION_DATE      = "expirationDate"
	CREDENTIAL_SIGNATURE = "signature"
	CONTEXT              = "context"
	CREDENTIAL_ID        = "id"
	ISSUANCE_DATE        = "issuanceDate"
	POLICY               = "Policy"
	POLICY_PACKAGE       = "com.webank.weid.cpt.policy."

	/**
	 * Person sonSchemaData.
	 */
	PERSON_DID  = "did"
	PERSON_TYPE = "type"
	PERSON_NAME = "name"
	PERSON_NUM  = "IDNumber"

	/**
	 * Person sonSchemaData.
	 */
	COMPANY_DID       = "did"
	COMPANY_TYPE      = "type"
	COMPANY_NAME      = "name"
	COMPANY_CODE      = "creditCode"
	COMPANY_LEGAL     = "legalName"
	COMPANY_LEGAL_DID = "legalDid"

	/**
	 * proof key.
	 */
	PROOF           = "proof"
	PROOF_SIGNATURE = "signatureValue"
	PROOF_TYPE      = "type"
	PROOF_CREATED   = "created"
	PROOF_CREATOR   = "creator"
	PROOF_SALT      = "salt"
)
