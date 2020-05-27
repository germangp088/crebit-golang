import React from 'react';

const Balance = (props) => {
    return (
      <div>
        <h2>{`Balance : $${props.balance.amount}`}</h2>
      </div>
    );
}

export default Balance;
