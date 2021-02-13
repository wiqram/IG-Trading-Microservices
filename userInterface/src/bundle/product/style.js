import styled from 'styled-components';
import { Row, Col } from '../common/style/styles';
import { colors } from '../common/style/base';

export const Card = styled(Row)`
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    box-shadow: 0 5px 5px 0 ${colors.shadowColor};
    border-top: 7px solid ${colors.primaryColor};
    margin-top: 5px;
    margin-left: 5px;
    margin-right: 5px;
    margin-bottom: 20px;
    padding-bottom: 0px;
    background-color: ${colors.cardColor};
    @media (max-width: 600px) and (min-width: 100px) {
      flex-direction: column;
    }
    `;
export const ProductName = styled.div`
    display: flex;
    align-items: flex-end;
    font-weight: 500;
    font-size: 28px;
    text-transform: uppercase;
    padding-right: 10px;
    color:${colors.primaryColor};
    word-wrap: break-word;;
    `;
export const CardRow = styled(Row)`
align-items: baseline;
    flex-direction: row;
    padding-top:10px;
    padding-bottom:15px;
    padding-left: 25px;
    @media (max-width: 600px) and (min-width: 100px) {
      flex-direction: column;
    }
    `;
export const AccountNumber = styled.div`
    display: flex;
    align-items: flex-end;
    font-weight: 300;
    font-size: 18px;
    color:#867777;
    justify-content: flex-start;
    `;
export const BalanceLabel = styled.div`
    display: flex;
    align-items: flex-end;
    font-weight: 300;
    font-size: 16px;
    color:#867777;
    padding-bottom: 2px;
    `;
export const Currency = styled.div`
    font-weight: 500;
    font-size: ${props => (props.size)}px;
     padding-right: 10px;
    `;
export const NavItems = styled(Col)`
    align-self: flex-end;
    `;