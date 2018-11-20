import smtplib
import datetime
from email.message import EmailMessage

from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText

from conf import loadConf
from output import rankingText, rankingHTML

CONF = loadConf()
USER = CONF['auth']['user']
PASSWORD = CONF['auth']['password']

def sendMail(weightedIssues):
    msg = MIMEMultipart('alternative')

    text = MIMEText(rankingText(weightedIssues), 'plain')
    html = MIMEText(rankingHTML(weightedIssues), 'html')

    msg.attach(text)
    msg.attach(html)

    msg['Subject'] = '{date} - RE/MAX Backlog Ranking : {n} issues'.format(n=len(weightedIssues), date=datetime.date.today())
    msg['From'] = ''
    msg['To'] = ''

    mail = smtplib.SMTP('smtp.gmail.com', 587)

    mail.ehlo()
    mail.starttls()
    mail.ehlo()

    mail.login(USER+'@homes.com', PASSWORD)
    mail.send_message(msg)
    mail.quit()
