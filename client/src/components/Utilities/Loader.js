import React from 'react';
import loader from './loader.module.scss';

export function RotatingCircleLoader(props) {
  const { className, height, width } = props;
  return (
    <div className={className}>
      <div
        className={`${loader.skCircle}`}
        style={{ height: height || 40, width: width || 40 }}
      >
        <div className={`${loader.skCircle1} ${loader.skChild}`} />
        <div className={`${loader.skCircle2} ${loader.skChild}`} />
        <div className={`${loader.skCircle3} ${loader.skChild}`} />
        <div className={`${loader.skCircle4} ${loader.skChild}`} />
        <div className={`${loader.skCircle5} ${loader.skChild}`} />
        <div className={`${loader.skCircle6} ${loader.skChild}`} />
        <div className={`${loader.skCircle7} ${loader.skChild}`} />
        <div className={`${loader.skCircle8} ${loader.skChild}`} />
        <div className={`${loader.skCircle9} ${loader.skChild}`} />
        <div className={`${loader.skCircle10} ${loader.skChild}`} />
        <div className={`${loader.skCircle11} ${loader.skChild}`} />
        <div className={`${loader.skCircle12} ${loader.skChild}`} />
      </div>
    </div>
  );
}
