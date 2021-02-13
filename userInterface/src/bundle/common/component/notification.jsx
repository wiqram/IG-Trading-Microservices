import React from 'react';
import { PropTypes } from 'prop-types';
import { Row, Alert } from "../style/styles";

class Notification extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        let message = this.props.message;

        return (
            <Alert>
                <Row>
                    <strong>{message}</strong>
                </Row>
            </Alert>
        );
    }
}
Notification.propTypes = {
    message: PropTypes.string.isRequired,
};
export default Notification;
