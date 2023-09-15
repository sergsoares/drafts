open System.Text.Json
open System.Text.Json.Serialization


type ColorEnum = Red=0 | Yellow=1 | Blue=2

printfn "%A" ColorEnum.Red

let yellow = ColorEnum.Yellow
printfn "%A" yellow

let test = enum<ColorEnum>(2)
printfn "%A" test


type StackStatus 
  = CREATE_COMPLETE 
  | CREATE_IN_PROGRESS 
  | CREATE_FAILED 
  | DELETE_COMPLETE
  | DELETE_FAILED
  | DELETE_IN_PROGRESS
  | REVIEW_IN_PROGRESS
  | ROLLBACK_COMPLETE
  | ROLLBACK_FAILED
  | ROLLBACK_IN_PROGRESS
  | UPDATE_COMPLETE
  | UPDATE_COMPLETE_CLEANUP_IN_PROGRESS
  | UPDATE_FAILED
  | UPDATE_IN_PROGRESS
  | UPDATE_ROLLBACK_COMPLETE
  | UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS
  | UPDATE_ROLLBACK_FAILED
  | UPDATE_ROLLBACK_IN_PROGRESS

printfn "%A" StackStatus.CREATE_COMPLETE

type Stack = {
  status: StackStatus
  name: string
  id: System.Guid
}

let jsonOptions = JsonFSharpOptions.Default().ToJsonSerializerOptions()

let mystack = { status = StackStatus.UPDATE_COMPLETE; name = "My personal stack"; id = System.Guid.NewGuid() }
printfn "%s" (JsonSerializer.Serialize(mystack, jsonOptions))  
