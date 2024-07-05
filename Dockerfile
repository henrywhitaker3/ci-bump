FROM scratch
WORKDIR /

COPY ci-bump /ci-bump

ENTRYPOINT ["/ci-bump"]
