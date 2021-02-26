import React, { useEffect } from 'react';
import { PropTypes } from 'prop-types';
import { useDispatch, useSelector } from 'react-redux';
import { Container } from '../../components/common/style/styles';
import ProductCard from './card';
import getProducts from '../../actions/product';
import Notification from '../../components/common/component/notification';
import Spinner from '../../components/common/component/spinner';
import {Breadcrumb, BreadcrumbItem} from "reactstrap";
/***
 * Display all products list here
 */

const ProductHome = () => {

    // set state for the props
    const useProduct = useSelector(state => ({
        products: state.productState.products,
        isLoading: state.productState.isLoading,
        error: state.productState.error
    }));

    // call dispatch for actions
    const dispatch = useDispatch();

    // call reducer action now
    useEffect(() => {
        dispatch(getProducts());
    }, []);

    // extract products list and others
    const { products, isLoading, error } = useProduct;

    if (isLoading === true) return <Spinner />;
    if (error) return <Notification message={error} />
    return (
        <Container id="productContainer">
            <Breadcrumb>
                <BreadcrumbItem>YOU ARE HERE</BreadcrumbItem>
                <BreadcrumbItem active>Products</BreadcrumbItem>
            </Breadcrumb>
            {
                products.map((product) => <ProductCard key={product.id} product={product} />)
            }
        </Container>
    )
}
ProductHome.propTypes = {
    products: PropTypes.object,
    isLoading: PropTypes.bool,
    error: PropTypes.string,
    dispatch: PropTypes.func,
};
export default ProductHome;