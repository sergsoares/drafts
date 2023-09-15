open System.Collections.Generic
open System
open k8s
open k8s.Models


let client = Environment.GetEnvironmentVariable("TEST")

printfn "%s" client

// [<EntryPoint>]
// let main argv =
//     let client =
//         Environment.GetEnvironmentVariable("KUBECONFIG")
//         |> KubernetesClientConfiguration.BuildConfigFromConfigFile
//         |> fun config -> new Kubernetes(config)

//     let command = (Array.ofList [ "sleep"; "3600" ]) :> IList<string>
//     let containers =
//         (Array.ofList
//             [ V1Container(image = "busybox", command = command, imagePullPolicy = "IfNotPresent", name = "busybox") ])
//         :> IList<V1Container>
//     let pod =
//         V1Pod
//             (apiVersion = "v1", kind = "Pod", metadata = V1ObjectMeta(Name = "busybox-fscharp"),
//              spec = V1PodSpec(containers = containers, restartPolicy = "Always"))
            
//     client.CreateNamespacedPodWithHttpMessagesAsync(pod, "default").Wait()
//     printfn "Pod deployed."
//     0 // return an integer exit code