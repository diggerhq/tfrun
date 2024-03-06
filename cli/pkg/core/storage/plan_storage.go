package storage

type PlanStorage interface {
	StorePlan(localPlanFilePath string, storedPlanFilePath string) error
	StorePlanFile(fileContents []byte, artifactName string, storedPlanFilePath string) error
	RetrievePlan(localPlanFilePath string, artifactName string, storedPlanFilePath string) (*string, error)
	DeleteStoredPlan(artifactName string, storedPlanFilePath string) error
	PlanExists(artifactName string, storedPlanFilePath string) (bool, error)
}
