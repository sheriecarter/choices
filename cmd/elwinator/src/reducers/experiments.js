import { sample } from '../shuffle';

const experimentInitialState = {
  id: '',
  namespace: '',
  name: '',
  segments: [],
  numSegments: 0,
  params: [],
};

const experiment = (state = experimentInitialState, action) => {
  switch(action.type) {
  case 'EXPERIMENT_ADD':
    return {
      ...state,
      id: action.id,
      namespace: action.namespace,
      name: action.name,
    };
  case 'EXPERIMENT_NAME':
    return {
      ...state,
      name: action.name,
    };
  case 'EXPERIMENT_NUM_SEGMENTS':
    const ns = parseInt(action.numSegments, 10);
    return {
      ...state,
      numSegments: ns,
      segments: sample(action.namespaceSegments, action.numSegments),
    };
  case 'PARAM_ADD':
    return {
      ...state,
      params: [...state.params, action.id],
    };
  case 'PARAM_DELETE':
    return {
      ...state,
      params: state.params.filter(id => id !== action.param),
    };
  default:
    return state;
  }
};

/**
 * getExperiment returns the experiment for the supplied id.
 * @param {Object} state - the experiments state object.
 * @param {string} experiment - the experiment id.
 */
export const getExperiment = (state, experiment) => {
  return state.find(e => e.id === experiment);
};

/**
 * getExperiments returns a list of experiment objects for the given experiment ids
 * @param {Object} state - the experiments state object.
 * @param {Array} experimentIDs - an array of experiment id's that you want returned.
 */
export const getExperiments = (state, experimentIDs) => {
  return experimentIDs.map(eid => state.find(e => eid === e.id));
};

/**
 * availableSegments returns the number of segments that are unused after combining the experiments.
 * @param {Object} state - the experiments state object.
 * @param {Array} experimentIDs - an array of experiment id's that you want returned.
 */
export const availableSegments = (state, experimentIDs) => {
  return getExperiments(state, experimentIDs).reduce((prev, e) => {
      return prev - e.numSegments;
    }, 128);
};

/**
 * combinedSegments returns the result of combining all segments given.
 * @param {Object} state - the experiments state object.
 * @param {Array} experimentIDs - an array of experiment id's that you want combined.
 */
export const combinedSegments = (state, experimentIDs) => {
  return getExperiments(state, experimentIDs)
  .reduce((prev, e) => {
    e.segments.forEach((seg, i) => {
      prev[i] |= seg
    });
    return prev;
  }, new Uint8Array(16).fill(0));
};

const experiments = (state = [], action) => {
  switch (action.type) {
  case 'EXPERIMENT_ADD':
    return [...state, experiment(undefined, action)];
  case 'EXPERIMENT_DELETE':
    return state.filter(e => e.id !== action.experiment);
  case 'EXPERIMENT_NAME':
  case 'EXPERIMENT_NUM_SEGMENTS':
  case 'PARAM_ADD':
  case 'PARAM_DELETE':
    const exps = state.map(e => {
      if (e.id !== action.experiment) {
        return e;
      }
      return experiment(e, action);
    });
    return exps;
  default:
    if (action.entities && action.entities.experiments) {
      return action.entities.experiments;
    }
    return state;
  }
};

export default experiments;
