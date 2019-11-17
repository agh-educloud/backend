package class_commons

var CodesToClassUuid = make(map[string]string)
var ClassUuidToCode = make(map[string]string)

// ClassID -> List[StudentID]
var AllStudentsInClass = make(map[string][]string)
