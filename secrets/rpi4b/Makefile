.PHONY: install
install:
	install -o root -g root -m 600 etc/postfix/sasl_passwd /etc/postfix/
	install -o root -g root -m 600 etc/nginx/letsencrypt/account.key /etc/nginx/letsencrypt/
	install -o root -g root -m 600 etc/nginx/letsencrypt/domain.csr /etc/nginx/letsencrypt/
	install -o root -g root -m 644 etc/nginx/letsencrypt/domain.key /etc/nginx/letsencrypt/
	install -o root -g root -m 644 etc/nginx/letsencrypt/dhparam.pem /etc/nginx/letsencrypt/
