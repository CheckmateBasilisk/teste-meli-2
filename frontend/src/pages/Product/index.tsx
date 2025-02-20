import './styles.scss'
import { useEffect, useState } from 'react'
import { useParams } from "react-router";

import Grid2 from '@mui/material/Grid2'
import Layout from '../../components/Layout/Layout'
import ProductInfo from './ProductInfo/ProductInfo';
import ProductActions from './ProductActions/ProductActions';

import { BaseUrl } from '../../utils/vars';

import itemImg from '/home/bagatini/Programming/teste-meli-2/frontend/src/assets/productImages/colorful-didactic-wooden-train-toy-for-preschool-aged-kids.webp'


const Product = () => {
    const {id} = useParams()
    const [item, setP] = useState({
        id: "",
        barcode: "",
        name: "",
        rating: 0,
        price: 0,
        stock: 0,
        descr: "",
        image: itemImg,
        alt: "", 
      })

    // fetch product info and put into item
    const url = new URL('/v1/product/'+id, BaseUrl).href
    useEffect(
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
        , []
    )
    
    return (
        <Layout>
                <Grid2 container spacing={1} size={9}> 
                    <Grid2 size={{ sm: 12, md: 4}}>
                        {/* product img */}
                        <div style={{ padding: '1rem', border: '1px solid #ccc', borderRadius: '0.5rem' }}>
                            <img src={item.image} alt={item.alt} style={{width:"100%",height:"auto" }}/>
                        </div>
                    </Grid2>
                    <Grid2 size={{ sm: 6, md: 5}}>
                        {/* product info */}
                        <ProductInfo product={item}/>
                    </Grid2>
                    <Grid2 size={{ sm: 6, md: 3}}>
                        {/* product actions (add to cart etc) */}
                        <ProductActions product={item}/>
                    </Grid2>
                </Grid2>
        </Layout>
    )
}

export default Product
