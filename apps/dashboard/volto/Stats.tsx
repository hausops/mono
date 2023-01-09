import * as s from './Stats.css';

type StatsProps = {
  value: string | number | undefined;
  fallbackValue?: string | number;
  unit: string;
};

export function Stats({value, fallbackValue = '-', unit}: StatsProps) {
  return (
    <div className={s.Stats}>
      <span className={s.DisplayValue}>{value ?? fallbackValue}</span>
      <span className={s.Unit}>{unit}</span>
    </div>
  );
}
