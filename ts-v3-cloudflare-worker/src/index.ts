import { serve } from 'inngest/cloudflare';
import { inngest } from './client';
import { fn1 } from './fns';

export default {
  fetch: serve({
    client: inngest,
    functions: [fn1],
  }),
};
