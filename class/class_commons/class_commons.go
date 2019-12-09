package class_commons

import web_gen "../../generated/protos/rest"

var CodesToClassUuid = make(map[string]string)
var ClassUuidToCode = make(map[string]string)

// ClassID -> List[StudentID]
var AllStudentsInClass = make(map[string][]string)

// TODO lock this array
var Statistics []web_gen.QuizQuestionStatistics

var PresentationHistoryData = make(map[int32][]web_gen.QuizQuestionStatistics)

var MessagesToPresenter []*web_gen.RestChatMessage
