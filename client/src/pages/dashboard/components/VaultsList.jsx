import VaultItem from "./VaultItem";

const VaultsList = ({ vaults, setVaults }) => {
  return (
    <div className="flex-grow flex items-center justify-center">
      {vaults.length === 0 ? (
        <p className="text-center text-gray-500 ">
          No vaults available. Try creating a new vault.
        </p>
      ) : (
        <div className="flex flex-col gap-4 w-full md:w-1/2 mt-4">
          {vaults.map((vault) => (
            <VaultItem key={vault.id} vault={vault} />
          ))}
        </div>
      )}
    </div>
  );
};

export default VaultsList;
