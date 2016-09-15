package models


type SessionData struct {
	IdUsuario string
}


func GetNewSession(idUsuario string) SessionData{
	return SessionData{IdUsuario: idUsuario}
}


type MessageReturn struct{
	Code int64
	Message string
}


func GetNewMessage(codep int64, message string) MessageReturn{
	return MessageReturn{Code: codep, Message: message}
}