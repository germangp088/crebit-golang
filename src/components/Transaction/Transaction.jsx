import React from 'react';
import {
    AccordionItem,
    AccordionItemHeading,
    AccordionItemButton,
    AccordionItemPanel
} from 'react-accessible-accordion';

const Transaction = (props) => {
    return (
        <AccordionItem>
            <AccordionItemHeading>
                <AccordionItemButton className={`accordion__button accordion__button--${props.type}`}>
                    {`Type: ${props.type} - Amount: $${props.amount}`}
                </AccordionItemButton>
            </AccordionItemHeading>
            <AccordionItemPanel>
                <div className="flex-container">
                    <div>{`Id : ${props.transaction_id}`}</div>
                    <div>{`Type : ${props.type}`}</div>
                    <div>{`Amount : $${props.amount}`}</div>
                    <div> {`Effective Date : ${props.effectiveDate}`}</div>
                </div>
            </AccordionItemPanel>
        </AccordionItem>
    )
}

export default Transaction
