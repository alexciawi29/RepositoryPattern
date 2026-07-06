import os
import re

go_files = [
    r'D:\AlexCiawi\mobile\ReposityPattern\Models\Vendor.go',
    r'D:\AlexCiawi\mobile\ReposityPattern\Controller\V6\MetadataController.go'
]

renames = {
    'RegistrationPurposes': 'RegistrationPurpose',
    'WorkOrderRemarks': 'WorkOrderRemark',
    'CreditTermDays': 'CreditTermDay',
    'DeliveryCityAreas': 'DeliveryCityArea',
    'MapsUrl': 'MapUrl',
    'Certifications': 'Certification',
}

for file in go_files:
    with open(file, 'r', encoding='utf-8') as f:
        content = f.read()
    
    for old, new in renames.items():
        content = content.replace(old, new)
    
    # Handle Details carefully
    content = content.replace('Details                string', 'Detail                 string')
    content = content.replace('json:"Details"', 'json:"Detail"')
    content = content.replace('Name="Details"', 'Name="Detail"')
    
    with open(file, 'w', encoding='utf-8') as f:
        f.write(content)

# Dart files
dart_dir = r'D:\AlexCiawi\mobile\vendor_portal\lib'
for root, _, files in os.walk(dart_dir):
    for file in files:
        if file.endswith('.dart'):
            path = os.path.join(root, file)
            with open(path, 'r', encoding='utf-8') as f:
                content = f.read()
            
            orig_content = content
            for old, new in renames.items():
                content = content.replace(old, new)
            
            # Handle Details carefully in Dart
            content = re.sub(r'\bDetails\b(?!View)(?!SectionWidget)', 'Detail', content)
            
            if content != orig_content:
                with open(path, 'w', encoding='utf-8') as f:
                    f.write(content)
print('Done rename')
