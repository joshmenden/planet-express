import React from 'react';
import Paper from '@material-ui/core/Paper';

const Ship = (props) => {
  const paperStyle = {
    overflowY: 'scroll',
    fontSize: 12,
    align: 'center',
    maxWidth: '80%',
    textAlign: 'left'
  }

  if (props.ship) {

    return(
      <Paper style={paperStyle}>
        <code>
          <pre>{JSON.stringify(props.ship, null, 2) }</pre>
        </code>
      </Paper>
    )
  } else {
    return(
      <div />
    )
  }
}

export default Ship