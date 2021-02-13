import styled from 'styled-components';
import { colors } from '../common/style/base';
import { Row, Button } from '../common/style/styles';

export const OtherLabel = styled.div`
    font-size: 14px;
    text-transform: uppercase;
    word-wrap: break-word;
    `;
export const DateLabel = styled.div`
    font-weight: 300;
    font-size: 16px;
    color:${colors.primaryColor};
    `;
export const TRow = styled(Row)`
    flex-direction: row;
    padding-top:12px;
    padding-bottom:12px;
    padding-left: 25px;
    background-color : ${props => props.color ? props.color : colors.rowColor};
    margin-bottom: 1px;
    @media (max-width: 600px) and (min-width: 100px) {
      flex-direction: column;
    }
    `;
export const Card = styled(Row)`
    display: flex;
    flex-direction: column;
    flex-wrap: wrap;
    box-shadow: 0 5px 5px 0 ${colors.shadowColor};
    border-top: 7px solid ${colors.primaryColor};
    margin-top: 5px;
    margin-left: 5px;
    margin-right: 5px;
    margin-bottom: 20px;
    padding: 20px;
    background-color: ${colors.cardColor};
    @media (max-width: 600px) and (min-width: 100px) {
      flex-direction: column;
    }
    `;
export const CardRow = styled(Row)`
    align-items: center;
    justify-content: center;
    @media (max-width: 600px) and (min-width: 100px) {
        flex-direction: column;
    }
    `;
export const FormRow = styled(Row)`
    display: flex;
    width: fill-available;
    margin-top: -10px;
    flex-wrap: wrap;
    padding-bottom: 15px;
    @media (max-width: 600px) and (min-width: 100px) {
        flex-direction: column;
    }
    `;
export const Label = styled.div`
    font-size: 16px;
    `;
export const ProductName = styled.div`
    font-size: 18px;
    text-transform: uppercase;
    color:${colors.primaryColor};
    word-wrap: break-word;;
    `;
export const AccountNumber = styled.div`
    display: flex;
    font-size: 16px;
    color:#867777;
    `;
export const TabButton = styled(Button)`
    margin-right: 2px;
    background-color : ${props => (props.isActive === props.tabName) ? colors.primaryColor : 'grey'};
    @media (max-width: 600px) and (min-width: 100px) {
      margin-bottom: 2px;
      margin-right: 0px;
      width: fill-available;
    }
    `;    