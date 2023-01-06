import {Address} from '@/services/address';
import {SingleFamilyProperty} from '@/services/property';
import {Section} from '@/volto/Section';
import {Entry, EntryList} from './EntryList';

type PropertyInfoProps = {
  property: SingleFamilyProperty;
};

export function PropertyInfo({property}: PropertyInfoProps) {
  const {address, unit} = property;
  const [street, region] = Address.from(address).format();
  return (
    <Section title="Property info">
      <EntryList>
        <Entry
          label="Address"
          value={
            <div>
              <p>{street}</p>
              <p>{region}</p>
            </div>
          }
        />
        {unit.bedrooms ? <Entry label="Beds" value={unit.bedrooms} /> : null}
        {unit.bathrooms ? <Entry label="Baths" value={unit.bathrooms} /> : null}
        {unit.size ? (
          <Entry
            label="Size"
            value={`${Intl.NumberFormat('en-US').format(unit.size)} sq.ft.`}
          />
        ) : null}
      </EntryList>
    </Section>
  );
}
