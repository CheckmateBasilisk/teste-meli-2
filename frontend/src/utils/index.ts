export const FormatBRL = (v: number) => {
    return Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(v)
}

// how to fix this?
export const SimpleFetch = (url: string, setP: React.Dispatch<React.SetStateAction<any>> ) => {
    // setP  is a function that recieves a parameter d of type S and return an S (?)
    ()=>{
    fetch(url)
        .then((response) => {
            return response.json() //FIXME: add error handling
        }).then((data) => {
            setP(data);
            //console.log(data)
            return
        })
    }
}
