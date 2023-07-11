import json
import PyPDF2
from hanlp import *
import re


def extract_pdf():
    with open(r'./data./docx2pdf/18.pdf', 'rb') as f:
        pdf_reader = PyPDF2.PdfReader(f)
        text = ''
        for page in pdf_reader.pages:
            text += page.extract_text()
        retext = text
        text = re.sub(r'\n', '', text)
        return text, retext


text, retext = extract_pdf()
print(text)
model = hanlp.load(hanlp.pretrained.mtl.UD_ONTONOTES_TOK_POS_LEM_FEA_NER_SRL_DEP_SDP_CON_XLMR_BASE)
result = model(text)
ner = result['ner']

NAME = ''
BASE = ''
SCHOOL = ''
def ner_pro(ner):
    for i in ner:
        if i[1] == 'PERSON':
            global NAME
            NAME = i[0]

            break

    for i in ner:
        if i[1] == 'ORG' and len(i[0]) > 4 and re.search(r'学', text):
            global SCHOOL
            SCHOOL = i[0]
            SCHOOL = re.sub(r' ', '', SCHOOL)
            break

    for a in ner:
        if a[1] == 'GPE':
            global BASE
            BASE = a[0]
            break
    return NAME,SCHOOL,BASE


name,school,base = ner_pro(ner)

email_pattern = r'\b[A-Za-z0-9._%+-]+ ?@ ?[A-Za-z0-9.-]+ ?\.[A-Za-z]{2,}\b'
phone_pattern = r'\b\d(?:\s*\d){10}\b'
birth_pattern = r"\b19\d{2}\b"
birth_2000 = r"\b200\d{1}"

emails = re.findall(email_pattern, retext)
emails = [email.replace(' ', '') for email in emails]
phone = re.findall(phone_pattern, retext)
phone = [phone.replace(' ', '') for phone in phone]
name = name.replace(' ', '')

global birth

birth = re.findall(birth_pattern, retext)
birth_2000 = re.findall(birth_2000, retext)


if emails == []:
    emails = ''
else:
    emails = emails[0]

if phone == []:
    phone = ''
else:
    phone = phone[0]


if birth == []:
    birth = 0
else:
    birth = birth[0]
    birth = int(birth)
    if birth > 1900:
        birth = 2023 - birth + 1

birth = int(birth)

if birth == 0:
    if birth_2000 == []:
        birth == 0
    else:
        birth = 2023-int(birth_2000[0])+1


education_pattern = r'(中专|大专|本科|硕士|博士)'
matches = re.findall(education_pattern, text)
highest_education = max(matches, key=lambda x: len(x))

work_year = 0

if highest_education == '本科':
    work_year = birth - 22
if highest_education == '硕士':
    work_year = birth - 25
if highest_education == '大专':
    work_year = birth - 20
if highest_education == '中专':
    work_year = birth - 18
if highest_education == '博士':
    work_year = birth - 28

if work_year<0:
    work_year = ''
if birth == 0:
    age = ''
keywords = ["求职目标", "求职意向"]

my_dict = {"highest_education": highest_education, "work_year": work_year, "base": BASE,
           "name": NAME, "school": SCHOOL,
           "email": emails, "phone": phone, "age": birth}
global job_t
job_t = ''
# 遍历关键词列表，查找匹配的关键词
for keyword in keywords:
    if keyword in retext:
        match = re.findall(r"(\S+)", text.split(keyword)[-1].strip("："))

        job_t = match[0]
        break
my_dict["job_target"] = job_t
json_str = json.dumps(my_dict, ensure_ascii=False)
print(json_str)