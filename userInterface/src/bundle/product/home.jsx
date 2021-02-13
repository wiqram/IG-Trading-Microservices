import React from 'react';
import { PropTypes } from 'prop-types';
import { connect } from 'react-redux';
import { Container } from '../common/style/styles';
import ProductCard from './card';
import getProducts from './middleware/actions';
import Notification from '../common/component/notification';
import Spinner from '../common/component/spinner';

@connect((state) => {
    return {
        products: state.productState.products,
        isLoading: state.productState.isLoading,
        error: state.productState.error
    };
})

class ProductHome extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            products: [],
            isLoading: false,
            error: ''
        };
    }
    componentDidMount() {
        this.props.dispatch(getProducts());
    }
    /** Call each time store dispatch called **/
    static getDerivedStateFromProps(nextProps) {
        return {
            products: nextProps.products,
            isLoading: nextProps.isLoading,
            error: nextProps.error
        };
    }
    render() {
        const { products } = this.props;
        if (this.state.isLoading == true) {
            return <Spinner />;
        } else if (this.state.error != null && this.state.error.trim().length > 0) {
            return <Notification message={this.state.error} />;
        }
        return (
            <Container id="productContainer">
                {
                    products.map((product) => <ProductCard key={product.id} product={product} />)
                }
            </Container>
        );
    }
}
ProductHome.propTypes = {
    products: PropTypes.object,
    isLoading: PropTypes.bool,
    error: PropTypes.string,
    dispatch: PropTypes.func,
};
export default ProductHome;