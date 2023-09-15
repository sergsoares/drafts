open System.Text.Json

type Content = {
  msg: string
}

type jwtHeader = {
  typ: string
  alg: string
  msg: Content
}

let msg2 = { msg = "teste"}
let exHeader = { typ = "JWT"; alg = "HS256"; msg = msg2 }
let exHeader2 = { typ = "JWT"; alg = "XPTO"; msg = msg2 }
let exHeader3 = { typ = "JWT"; alg = "New"; msg = msg2 }

let allTypes = ResizeArray<jwtHeader>()
allTypes.Add exHeader 
allTypes.Add exHeader2 
allTypes.Add exHeader3

printfn "%s" (JsonSerializer.Serialize exHeader)
printfn "%s" (JsonSerializer.Serialize allTypes)

