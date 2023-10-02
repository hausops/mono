import {Button} from '@/volto/Button';
import {EmptyState} from '@/volto/EmptyState';
import {HomeIcon} from '@/volto/icons';
import Link from 'next/link';
import * as s from './PropertyNotFound.css';

export function PropertyNotFound() {
  return (
    <div className={s.NotFound}>
      <EmptyState
        icon={<HomeIcon />}
        title="Property not found"
        description="Get started by adding a new property."
        actions={
          <>
            <Button as={Link} variant="outlined" href="/properties">
              All properties
            </Button>
            <Button as={Link} variant="contained" href="/properties/new">
              Add property
            </Button>
          </>
        }
      />
    </div>
  );
}
