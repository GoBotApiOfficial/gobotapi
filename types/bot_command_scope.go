package types

// BotCommandScope Represents the scope to which bot commands are applied
// Currently, the following 7 scopes are supported:
//  - BotCommandScopeDefault
//  - BotCommandScopeAllPrivateChats
//  - BotCommandScopeAllGroupChats
//  - BotCommandScopeAllChatAdministrators
//  - BotCommandScopeChat
//  - BotCommandScopeChatAdministrators
//  - BotCommandScopeChatMember
type BotCommandScope interface{}
