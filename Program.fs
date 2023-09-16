//Copy this code on your program file

let wordfilter a b=
    
    let newList = []
    for i in a do
        let h = string i
        let x = string b
        let subwords = h.Split([|x|], System.StringSplitOptions.None)
        let newWords = List.ofArray subwords
        let length = List.length newWords
        if length > 1 then
            let newList = h :: newList
            printfn "Phrase %s" i
            ""
        else
            "nothing"

let mapNthElement f n list =
    let rec mapNthElement' f count acc = function
        | [] -> List.rev acc
        | x :: xs when count < n ->
            let result = f x
            mapNthElement' f (count + 1) (result :: acc) xs
        | _ :: _ -> List.rev acc

    mapNthElement' f 0 [] list


let main() =
    printfn "This is the code for the word filter"

    wordfilter ["You're such a crybaby"; "This is such a great song"; "Let's go to the Mall"] "such"
    
    let myList = [12; 1342; 32; 49; 35]
    let n = 2
    let result = mapNthElement id n myList
    let element = List.item (n-1) result

    printfn "the number in the %d position is %d" n element



// For more information see https://aka.ms/fsharp-console-apps
main()
