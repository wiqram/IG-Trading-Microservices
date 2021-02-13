import styled from 'styled-components';
import { colors } from './base';

export const Container = styled.div`
    /* width: 100%; */
    margin: -5px;
    padding: 5px;
    background-color: ${colors.background}
    `;

export const Col = styled.div`
    display: flex;
    flex-direction: column;
    flex: ${props => (props.col ? props.col : 1)};
    text-align: ${props => props.textAlign};
    align-items: ${props => props.alignItems};
    width: fill-available;
    `;
export const Row = styled.div`
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    padding-top: 5px;
    padding-bottom: 5px;
    align-items: baseline;
    `;

export const Button = styled.button`
    width: fill-available;
    cursor: pointer;
    outline:0 !important; 
    font-size: 16px;
    font-weight: 400;
    color: #ffffff;
    border: 1px solid transparent;
    text-align: center;
    height: 45px;
    `;

export const ButtonPrimary = styled(Button)`
    background-color: ${colors.primaryColor};
    `;
export const ButtonSuccess = styled(Button)`
    background-color: #5cb85c;
    `;
export const ButtonInfo = styled(Button)`
    background-color: #5bc0de;
    `;
export const ButtonDanger = styled(Button)`
    background-color: #d9534f;
    `;
export const FormDiv = styled.div`
    opacity: 1;
    display: block;
    display: block;
    width: 100%;
    padding-top: 0.375rem;
    padding-bottom: 0.375rem;
    line-height: 1.5;
    background-color: #ebf1f7;
    `;
export const SubmitButton = styled.button`
    background-color: #012d65;
    display: inline-block;
    font-weight: 500;
    color: #ffffff;
    border: 1px solid transparent;
    padding: 10px 10px;
    line-height: 1.5;
    border-radius: 0.25rem;
    text-align: center;
    `;
export const Input = styled.input`
    padding: 0.5em;
    margin: 0.5em;
    background: #bdbdbd;
    border: none;
    border-radius: 3px;
    width: auto;
    height: 25px;
`;

export const Alert = styled.div`
    border-radius: 0.25rem;
    width: 100%;
    padding: 20px;
    background-color: #5bc0de;
`;
