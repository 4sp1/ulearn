# SMTP Server

> ⚠︎ This is Work in Progress

Here's a documented, reproducible process for setting up a Dovel email server:

---

## **Email Server Setup with Dovel**

### **Prerequisites**
- A deployed VM with a public IP address
- Domain name with DNS management access
- Root/sudo access to the VM
- `DOVEL_CONFIG_DIR` is set (you should use `/home/dovel/.config/dovel`)

---

### **Step 1: Initial Server Setup**

**1.1 Create a dedicated system user**
```bash
# Create user without login shell
sudo useradd -r -s /bin/false -m -d /var/lib/dovel dovel

# Create config directory
sudo mkdir -p $DOVEL_CONFIG_DIR
sudo chown dovel:dovel /home/.config/dovel
```

**1.2 Configure firewall (if needed)**
```bash
# Allow email ports
sudo ufw allow 25/tcp   # SMTP
sudo ufw allow 587/tcp  # Submission
sudo ufw allow 465/tcp  # SMTPS
sudo ufw allow 993/tcp  # IMAPS
sudo ufw allow 8000/tcp # Admin/API
```

---

### **Step 2: Generate DKIM Keys**

**2.1 Create and secure key directory**
```bash
sudo -u dovel mkdir -p $DOVEL_CONFIG_DIR/keys
```

**2.2 Generate 3072-bit RSA key**
```bash
cd $DOVEL_CONFIG_DIR/keys
sudo -u dovel openssl genrsa -out dkim_private.pem 3072
sudo -u dovel chmod 600 dkim_private.pem
```

**2.3 Extract public key for DNS**
```bash
# Extract and format public key (removing headers/footers and line breaks)
PUBLIC_KEY=$(sudo -u dovel openssl rsa -in dkim_private.pem -pubout 2>/dev/null | \
    sed -e '/-----/d' | tr -d '\n')
echo "TXT DKSIM Record Value For DNS:"
echo "v=DKIM1; k=rsa; p=$PUBLIC_KEY"
```

---

### **Step 3: Configure DNS Records**

Add these records to your domain's DNS:

**3.1 MX Record** (Mail Exchange):
```
Type: MX
Name: @ (or subdomain)
Value: mail.yourdomain.com (your VM's hostname)
Priority: 10
```

> ⚠︎ semantic of "Name" need to be precised, further inspection is needed.

**3.2 A Record** for mail server:
```
Type: A
Name: mail.yourdomain.com
Value: [YOUR_VM_PUBLIC_IP]
```

**3.3 DKIM TXT Record**:
```
Type: TXT
Name: default._domainkey
Value: "v=DKIM1; k=rsa; p=[PUBLIC_KEY_HERE]"
```

**3.4 Optional: SPF Record**:
```
Type: TXT
Name: @
Value: "v=spf1 mx -all"
```

> ⚠︎ regarding records, we should learn more about this,
> at the moment we are unable to receive sent email

---

### **Step 4: Install Dovel Server**

**4.1 Install Go (if not present)**
```bash
sudo -u dovel 'cd && curl -o go.tgz -L -s https://go.dev/dl/go1.25.5.linux-amd64.tar.gz'
sudo 'rm -rf /usr/local/go && tar -zxf go.tgz -C /usr/local'
sudo -u dovel 'cd && rm go.tgz'
```

**4.2 Install Dovel server**
```bash
sudo -u dovel go install dovel.email/server@v0.13.1
```

**4.3 Locate the binary** (usually in `~/go/bin` for the dovel user)

---

### **Step 5: Create Configuration Files**

**5.1 Create `config.json`** at `$DOVEL_CONFIG_DIR/config.json`:
```bash
tee $DOVEL_CONFIG_DIR/config.json <<EOT
{
    "Port": "25",
    "Domain": "mail.yourdomain.com",
    "VaultFile": "/home/.config/dovel/users.json",
    "KeyFile": "/etc/dovel/keys/dkim_private.pem"
}
```

**5.2 Create `users.json`** at `$DOVEL_CONFIG_DIR/users.json`:
```json
tee $DOVEL_CONFIG_DIR/users.json <<EOT
[
    {
        "Name": "user@mail.yourdomain.com",
        "PrivateKey": "$DOVEL_CONFIG_DIR/keys/dkim_private.pem",
        "Password": "user_very_secure_password"
    }
]
EOT
```

**5.3 Set proper permissions:**
```bash
sudo chown dovel:dovel $DOVEL_CONFIG_DIR/{config,users}.json
sudo chmod 600 $DOVEL_CONFIG_DIR/users.json
```

### **Step 6: Testing & Verification**

**6.1 Test SMTP connectivity:**
```bash
nc mail.yourdomain.com 25
```

You might need to install `netcat`, for example
```bash
sudo apt install netcat-openbsd
```

**6.2 Verify DNS records:**
```bash
# Check MX record
dig MX yourdomain.com +short

# Check DKIM record
dig TXT default._domainkey.yourdomain.com +short
```

**7.3 Send test email** using mail client configured with:
- SMTP: mail.yourdomain.com:587 (STARTTLS) or :465 (SSL)
- Username: user@mail.yourdomain.com
- Password: your_secure_password_here

> We should add dig proper installation notes

---

### **Troubleshooting Notes**

1. **Port 25 issues**: Some cloud providers block port 25 by default. You may need to:
   - Request port 25 unblocking from your provider
   - Use ports 587/465 and configure Dovel accordingly
   - Adjust firewall rules

2. **DKIM validation**: Test with:
   ```bash
   sudo -u dovel opendkim-testkey -d yourdomain.com -s default -k /etc/dovel/keys/dkim_private.pem -vvv
   ```

---

### **Security Recommendations**

1. **Regular updates**:
   ```bash
   sudo -u dovel go get -u dovel.email/server@latest
   ```

2. **Monitor logs**: for instance use `fail2ban`

3. **SSL/TLS**: Consider adding SSL certificates for encrypted connections

4. **Backup DKIM keys** and configuration files
