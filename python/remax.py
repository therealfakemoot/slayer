from conf import setupFlags

import arrow

from issue import Issue
from board import Board,Backlog
from output import rankingHTML
from send import sendMail

def U(issue):
    pMap = {
        "blocker": 10,
        "critical": 8,
        "major": 6,
        "minor": 3,
        "trivial": 1
    }
    return (issue.SLA  + pMap[issue.priority.lower()]) ** 1.25

def Rank(bl, keyFilter):
    ret = []
    for i in sorted(bl, key=U, reverse=True):
        if i.key.startswith(keyFilter):
            ret.append({'issue':i, 'weight':U(i)})

    return ret

def main(args):
    b = Board(args.board)

    fields = {'fields':['priority,status,labels,components,created,summary,description']}
    bl = Backlog(b, fields)
    sendMail(Rank(bl, args.filter))

if __name__ == '__main__':
    p = setupFlags()

    args = p.parse_args()
    if not args.board:
        p.print_help(sys.stderr)
        sys.exit(0)

    main(args)
