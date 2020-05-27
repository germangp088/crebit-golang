import React from 'react';
import './Transactions.css';
import { Accordion } from 'react-accessible-accordion';
import Transaction from './Transaction';

class Transactions extends React.Component {

  render() {
      const accordionItems = this.props.transactions.map(transaction => {
        const { transaction_id, type, amount, effectiveDate } = transaction

        return (
            <Transaction
              key={transaction_id}
              transaction_id={transaction_id}
              type={type}
              amount={amount}
              effectiveDate={effectiveDate} />
        )
    })
    return (
      <Accordion allowMultipleExpanded={true} allowZeroExpanded ={true}>
        {accordionItems.length === 0 ? <div>No transactions to show, check README.md instruction to create them</div> : 
        <div>
          {accordionItems}
        </div> }
      </Accordion>
    );
  }
}

export default Transactions;
