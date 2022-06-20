import React, { useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import FamilyForm from '../react-hook-forms/FamilyForm';
import styles from './Counter.module.css';

export function Counter() {
  const dispatch = useDispatch();
  const [incrementAmount, setIncrementAmount] = useState('2');

  const incrementValue = Number(incrementAmount) || 0;

  return (
    <div>
      <div className={styles.row}>
        <FamilyForm />
      </div>
      <hr></hr>
      <div className={styles.row}>
        <button>Add Async</button>
        <button>Add If Odd</button>
      </div>
    </div>
  );
}
