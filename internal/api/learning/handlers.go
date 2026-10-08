package learning

import (
	"github.com/ArteShow/DoThat/internal/learning"
)

type LearningHandler struct {
	LearningManager *learning.LearningManager
}

func NewLearningHandler(learningManager *learning.LearningManager) *LearningHandler {
	return &LearningHandler{
		LearningManager: learningManager,
	}
}
