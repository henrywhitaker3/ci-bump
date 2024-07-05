FROM scratch
WORKDIR /

COPY ci-bump /ci-bump
USER 65532:65532

ENTRYPOINT ["/ci-bump"]
