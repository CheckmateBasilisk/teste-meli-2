import './styles.scss'
import { useEffect, useState } from 'react'
import ProductCard from '../../components/ProductCard/ProductCard'

import Grid2 from '@mui/material/Grid2'
import Divider from '@mui/material/Divider';
import Layout from '../../components/Layout/Layout'

import itemImg from '/home/bagatini/Programming/teste-meli-2/frontend/src/assets/productImages/colorful-didactic-wooden-train-toy-for-preschool-aged-kids.webp'
import { useSearchParams } from 'react-router';
import { BaseUrl } from '../../utils/vars';

const Search = () => {
    const[search, setSearch] = useState("")
    const [products, setProducts] = useState([{
        id: "",
        barcode: "",
        name: "",
        rating: 0,
        price: 0,
        stock: 0,
        descr: "",
        image: itemImg,
        alt: "", 
      }]
    )

    // fetch products info and put into products
    // FIXME: for now it fetches all products
    const [searchParams, setSearchParams] = useSearchParams();
    let productName = searchParams.get('name')
    const url = new URL('/v1/product/search?name='+productName, BaseUrl).href
    useEffect(
        ()=>{
            fetch(url)
                .then((response) => {
                    return response.json() //FIXME: add error handling
                }).then((data) => {
                    setProducts(data);
                    return
                })
        }
        , [search]
    )

    return (
        <Layout>
            <Grid2 container spacing={2}>
                <Grid2 container size={12}>
                    {products.map((product,key) => (
                        <Grid2 size={4} key={key}>
                            <ProductCard product={product}/>
                        </Grid2>
                    ))}
                </Grid2>
            </Grid2>
        </Layout>
    )
}

export default Search
