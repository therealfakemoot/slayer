from functools import total_ordering

import arrow

@total_ordering
class Issue():
    def __init__(self, raw):
        self.key = raw['key']
        self.fields = raw['fields']

        self.priority = self.fields['priority']['name']
        self.summary = self.fields['summary']
        self.components = [c.get('name') for c in self.fields['components']]
        self.created = arrow.get(self.fields['created']).datetime
        self.description = self.fields['description']
        self.status = self.fields['status']
        self.labels = self.fields['labels']

    @property
    def age(self):
        '''Returns the number of days ticket has not met SLA.'''
        today = arrow.get().date()

        delta = ((today - self.created.date()).days)
        return delta

    @property
    def SLA(self):
        '''Returns the number of days ticket has not met SLA.'''
        today = arrow.get().date()

        delta = ((today - self.created.date()).days) - 5

        if delta <= 0:
            return 0
        return delta

    def __lt__(self, other):
        return self.created < other.created

    def __eq__(self, other):
        return self.key == other.key

    def __repr__(self):
        return '<<{key}> <Priority:{priority}> <Summary {summary}> <Created {created}>>'.format(key=self.key, priority=self.priority, summary=self.summary, created=self.created)
