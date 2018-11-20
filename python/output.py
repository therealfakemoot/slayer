from string import Template
import json

import jinja2

def rankingText(weightedIssues):
    lines = []
    for weighted in weightedIssues:
        i = weighted['issue']
        lines.append('Key: {key}    Priority: {priority}    Summary:{summary}'.format(key=i.key, priority=i.priority, summary=i.summary))
    return "\n".join(lines)

def rankingHTML(weightedIssues):
    template = '''
<table style="color: white;border: 1px solid black; font-weight: bold;">
    <tr style="background: black;" align="center" >
        <th>Key</th>
        <th>Priority</th>
        <th>Age</th>
        <th>Summary</th>
    </tr>
    {% for issue in issues %}
    {% set i = issue['issue'] %}
    <tr {% if i.SLA > 0 %}bgcolor="darkred"{% else %}style="background: orange;color: black;"{% endif %}>
        <td>{{ i.key }}</td>
        <td>{{ i.priority }}</td>
        <td>{{ i.age }} days</td>
        <td>{{ i.summary }}</td>
    </tr>
    {% endfor %}
</table>
'''

    T = jinja2.Template(template)
    return T.render(issues=weightedIssues)

def rankingJSON(weightedIssues):
    return json.dumps(weightedIssues)
