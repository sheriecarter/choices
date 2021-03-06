// @flow
import React, { PropTypes } from 'react';
import { Link } from 'react-router';
import { connect } from 'react-redux';

import { paramDelete } from '../actions';
import { paramURL } from '../urls';

const ParamList = ({ namespace, experimentID, params, dispatch }) => (
  <table className="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>Param</th>
        <th>Choices</th>
        <th>Delete</th>
      </tr>
    </thead>
    <tbody>
    {params.map((param, i) => 
      <tr key={param.name} >
        <td>{i + 1}</td>
        <td><Link to={paramURL(param.id)}>{param.name}</Link></td>
        <td>{param.choices.join(', ')}</td>
        <td><button className="btn btn-default btn-xs" onClick={
          () => dispatch(paramDelete(namespace, experimentID, param.id))
        }>&times;</button></td>
      </tr>
    )}
    </tbody>
  </table>
);

ParamList.propTypes = {
  experimentID: PropTypes.string.isRequired,
  params: PropTypes.arrayOf(PropTypes.object).isRequired,
  dispatch: PropTypes.func.isRequired,
}

const connected = connect()(ParamList);

export default connected;
