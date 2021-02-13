import { combineReducers } from 'redux';
import registerReducer from "../middleware/register/reducer"

const rootReducer = combineReducers(
    {
        registerState: registerReducer
    }
);

export default rootReducer;
